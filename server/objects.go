package server

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
)

var ENDPOINTS = []string{"192.168.79.134:2379", "192.168.79.136:2379", "192.168.79.137:2379"}

type ServiceRegister struct {
	Cli       *clientv3.Client
	LeaseID   clientv3.LeaseID
	KeepAlive <-chan *clientv3.LeaseKeepAliveResponse
	Key       string
	Value     string
}

func (r *ServiceRegister) PutKeyWithLease(lease int64) error {
	leaseGrantRes, err := r.Cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	_, err = r.Cli.Put(context.Background(), r.Key, r.Value, clientv3.WithLease(leaseGrantRes.ID))
	if err != nil {
		return err
	}
	keepAliveChan, err := r.Cli.KeepAlive(context.Background(), leaseGrantRes.ID)
	if err != nil {
		return err
	}
	r.LeaseID = leaseGrantRes.ID
	r.KeepAlive = keepAliveChan
	log.Printf("put key: %s, value: %s success\n", r.Key, r.Value)
	return nil
}

func (r *ServiceRegister) ListenKeepChan() {
	for keepLeaseChan := range r.KeepAlive {
		log.Printf("lease %s success:%s\n", r.Key, keepLeaseChan)
	}
	log.Printf("lease stopped %s\n", r.Key)
}

func (r *ServiceRegister) Close() error {
	_, err := r.Cli.Revoke(context.Background(), r.LeaseID)
	if err != nil {
		return err
	}
	log.Printf("lease removed for %s", r.Key)
	return r.Cli.Close()
}
