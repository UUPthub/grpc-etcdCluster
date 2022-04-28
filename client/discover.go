package client

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Register interface {
	WatchServer(prefix string)
	GetServerList() []string
}

func NewServiceDiscover() (Register, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   ENDPOINTS,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &ServiceDiscovery{
		Cli:       client,
		ServerMap: make(map[string]string),
	}, nil
}
