package etcdv2

import (
	"time"

	clientv2 "go.etcd.io/etcd/client/v2"
	"golang.org/x/net/context"
)

func Connect(endpoints []string) (*clientv2.Client, error) {
	cfg := clientv2.Config{
		Endpoints:               endpoints,
		Transport:               clientv2.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	}

	return clientv2.New(cfg)
}

func Set(kapi clientv2.KeysAPI, key, value string) error {
	_, err := kapi.Set(context.Background(), key, value, nil)
	return err
}

func Get(kapi clientv2.KeysAPI, key string) (string, error) {
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

func Del(kapi clientv2.KeysAPI, key string) error {
	_, err := kapi.Delete(context.Background(), key, nil)
	return err
}
