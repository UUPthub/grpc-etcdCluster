package server

import (
	"go.etcd.io/etcd/client/v3"
	"time"
)

type Register interface {
	PutKeyWithLease(lease int64) error
	ListenKeepChan()
	Close() error
}

func NewServiceRegister(key, value string, lease int64) (Register, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   ENDPOINTS,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	srv := &ServiceRegister{
		Cli:   client,
		Key:   key,
		Value: value,
	}
	err = srv.PutKeyWithLease(lease)
	if err != nil {
		return nil, err
	}
	return srv, nil
}
