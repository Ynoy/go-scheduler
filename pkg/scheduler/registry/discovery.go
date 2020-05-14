package registry

import (
	client "github.com/coreos/etcd/clientv3"
	"go-scheduler/pkg/web/conf"
	"go-scheduler/pkg/web/service"
	"log"
	"sync"
	"time"
)

type Service struct {
	instance *service.Instance
	leaseID  client.LeaseID
	close    chan struct{}
	swg      sync.WaitGroup
}

var (
	err    error
	Client *client.Client
)

// New Etcd V3 Client
func NewClient() {
	if Client, err = client.New(client.Config{
		Endpoints:   conf.Conf.Etcd.Endpoint,
		DialTimeout: 10 * time.Second,
	}); err != nil {
		log.Panicln(err)
	}
}
