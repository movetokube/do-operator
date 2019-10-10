package dns

import (
	"github.com/movetokube/do-operator/pkg/apis/do/v1alpha1"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"testing"
)

type DNSControllerTestSuite struct {
	suite.Suite
}

func (s *DNSControllerTestSuite) Test_IgnoresActiveCRs() {
	var (
		namespace = "sandbox"
		name      = "dns"
	)
	dodns := &v1alpha1.DNS{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.DNSSpec{
			DomainName: "example.com",
			//RecordType: "A",
			Hostname:   "www",
			Value: v1alpha1.DNSRecordValue{
				Literal: "127.0.0.1",
			},
		},
		Status: v1alpha1.DNSStatus{
			State: v1alpha1.STATE_ACTIVE,
			ID:    999,
		},
	}
	objs := []runtime.Object{dodns}

	sh := scheme.Scheme
	sh.AddKnownTypes(v1alpha1.SchemeGroupVersion, dodns)

	cl := fake.NewFakeClient(objs...)
	r := &ReconcileDNS{client: cl, scheme: sh}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: namespace,
			Name:      name,
		},
	}
	res, err := r.Reconcile(req)
	s.NoError(err)
	s.Assert().False(res.Requeue)

}

func TestDNSControllerSuite(t *testing.T) {
	suite.Run(t, new(DNSControllerTestSuite))
}