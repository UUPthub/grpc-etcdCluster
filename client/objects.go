package client

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
)

var ENDPOINTS = []string{"192.168.79.134:2379", "192.168.79.136:2379", "192.168.79.137:2379"}

type ServiceDiscovery struct {
	Cli       *clientv3.Client
	ServerMap map[string]string
	Lock      sync.Mutex
}

func (d *ServiceDiscovery) WatchServer(prefix string) {
	getRes, err := d.Cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Printf("etcd get error: %s\n", err.Error())
	}
	for _, kv := range getRes.Kvs {
		d.setServerMap(string(kv.Key), string(kv.Value))
	}
	go d.watcher(prefix)
}

func (d *ServiceDiscovery) setServerMap(k, v string) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	d.ServerMap[k] = v
	log.Printf("add server %s in server list\n", k)
}

func (d *ServiceDiscovery) delServerMap(k string) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	delete(d.ServerMap, k)
	log.Printf("delete server %s from server list\n", k)
}

func (d *ServiceDiscovery) GetServerList() []string {
	var service []string
	d.Lock.Lock()
	defer d.Lock.Unlock()
	for _, ser := range d.ServerMap {
		service = append(service, ser)
	}
	return service
}

func (d *ServiceDiscovery) watcher(prefix string) {
	watchChan := d.Cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching server %s\n", prefix)
	for resp := range watchChan {
		for _, kv := range resp.Events {
			switch kv.Type {
			case mvccpb.PUT:
				d.setServerMap(string(kv.Kv.Key), string(kv.Kv.Value))
			case mvccpb.DELETE:
				d.delServerMap(string(kv.Kv.Key))
			}
		}
	}
}
