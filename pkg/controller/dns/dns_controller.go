package dns

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/go-logr/logr"
	dov1alpha1 "github.com/movetokube/do-operator/pkg/apis/do/v1alpha1"
	"github.com/movetokube/do-operator/pkg/config"
	"github.com/movetokube/do-operator/pkg/digitalocean"
	"github.com/movetokube/do-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

var log = logf.Log.WithName("controller_dns")

// Add creates a new DNS Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileDNS{
		client:     mgr.GetClient(),
		scheme:     mgr.GetScheme(),
		dnsManager: digitalocean.NewDNSManager(config.GetConfig().DOToken),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("dns-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource DNS
	err = c.Watch(&source.Kind{Type: &dov1alpha1.DNS{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileDNS implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileDNS{}

// ReconcileDNS reconciles a DNS object
type ReconcileDNS struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client     client.Client
	scheme     *runtime.Scheme
	dnsManager digitalocean.DNSManager
}

// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileDNS) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling DNS Resource")

	// Fetch the DNS instance
	instance := &dov1alpha1.DNS{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	meta := instance.TypeMeta

	if instance.GetDeletionTimestamp() != nil {
		if instance.Status.State == dov1alpha1.STATE_ACTIVE {
			instance.Status.State = dov1alpha1.STATE_DELETING
			err = r.client.Status().Update(context.TODO(), instance)
			if err != nil {
				return reconcile.Result{}, err
			}
			statusCode, err := r.dnsManager.DeleteRecord(&digitalocean.RecordDeleteRequest{
				DomainName: instance.Spec.DomainName,
				RecordId:   instance.Status.ID,
			})
			if err != nil && statusCode != 404 {
				return reconcile.Result{
					RequeueAfter: time.Second * 5,
				}, err
			} else {
				err = r.removeFinalizer(reqLogger, instance, &meta)
				if err != nil {
					return reconcile.Result{
						RequeueAfter: time.Second * 5,
					}, err
				}
				return reconcile.Result{}, nil
			}
		} else {
			err = r.removeFinalizer(reqLogger, instance, &meta)
			if err != nil {
				return reconcile.Result{
					RequeueAfter: time.Second * 5,
				}, err
			}
			return reconcile.Result{}, nil
		}
	}

	if instance.Status.State == dov1alpha1.STATE_INITIAL {
		// set initial status first
		instance.Status.State = dov1alpha1.STATE_PENDING
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			return reconcile.Result{}, err
		}
		i := &instance.Spec
		id, err := r.dnsManager.CreateRecord(&digitalocean.RecordCreateRequest{
			DomainName: instance.Spec.DomainName,
			DomainRecordEditRequest: godo.DomainRecordEditRequest{
				Type:     i.RecordType,
				Name:     i.Hostname,
				Data:     i.Value.Literal,
				Priority: utils.ZeroIntIfNil(i.Priority),
				Port:     utils.ZeroIntIfNil(i.Port),
				TTL:      utils.ZeroIntIfNil(i.TTL),
				Weight:   utils.ZeroIntIfNil(i.Weight),
				Flags:    utils.ZeroIntIfNil(i.Flag),
				Tag:      utils.ZeroStringIfNil(i.Tag),
			},
		})
		if err != nil {
			return reconcile.Result{}, err
		}
		instance.Status = dov1alpha1.DNSStatus{
			State: dov1alpha1.STATE_ACTIVE,
			ID:    *id,
		}
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	err = r.addFinalizer(reqLogger, instance, &meta)
	if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("reconciled, doing nothing")
	return reconcile.Result{}, nil
}

func (r *ReconcileDNS) addFinalizer(reqLogger logr.Logger, i *dov1alpha1.DNS, meta *v1.TypeMeta) error {
	if len(i.GetFinalizers()) == 0 && i.GetDeletionTimestamp() == nil {
		reqLogger.Info("adding Finalizer")
		i.SetFinalizers([]string{getFinalizerName(meta)})

		// Update CR
		err := r.client.Update(context.TODO(), i)
		if err != nil {
			reqLogger.Error(err, "failed to update CR with finalizer")
			return err
		}
	}
	return nil
}

func getFinalizerName(i *v1.TypeMeta) string {
	return "finalizer." + i.APIVersion
}

func (r *ReconcileDNS) removeFinalizer(reqLogger logr.Logger, i *dov1alpha1.DNS, meta *v1.TypeMeta) error {
	finName := getFinalizerName(meta)
	for idx, fin := range i.GetFinalizers() {
		if fin == finName {
			finalizers := i.GetFinalizers()
			finalizers[idx] = finalizers[len(finalizers)-1]
			finalizers = finalizers[:len(finalizers)-1]
			i.SetFinalizers(finalizers)
			err := r.client.Update(context.TODO(), i)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
