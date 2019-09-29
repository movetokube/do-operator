package digitalocean

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/go-logr/logr"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"strconv"
)

type DNSManager interface {
	CreateRecord(req *RecordCreateRequest) (*int, error)
	DeleteRecord(req *RecordDeleteRequest) (int, error)
}

type RecordCreateRequest struct {
	DomainName string
	godo.DomainRecordEditRequest
}

type RecordDeleteRequest struct {
	DomainName string
	RecordId   int
}

type dnsMgr struct {
	client *godo.Client
	logger logr.Logger
}

func NewDNSManager(token string) DNSManager {
	return &dnsMgr{
		client: newClient(token),
		logger: logf.Log.WithName("DNSManager"),
	}
}

func (d *dnsMgr) CreateRecord(req *RecordCreateRequest) (*int, error) {
	record, _, err := d.client.Domains.CreateRecord(context.TODO(), req.DomainName, &req.DomainRecordEditRequest)
	if err != nil {
		d.logger.Error(err, "failed to create record")
		return nil, err
	} else {
		d.logger.Info("record created", "name", req.Name, "id", record.ID)
	}
	return &record.ID, err
}

func (d *dnsMgr) DeleteRecord(req *RecordDeleteRequest) (int, error) {
	res, err := d.client.Domains.DeleteRecord(context.TODO(), req.DomainName, req.RecordId)

	if err != nil {
		d.logger.Error(err, "failed to delete record")
		return res.StatusCode, err
	} else {
		d.logger.Info("record is deleted", "record Id", strconv.Itoa(req.RecordId))
	}
	return res.StatusCode, nil
}
