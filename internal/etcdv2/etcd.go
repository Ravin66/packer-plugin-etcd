package etcdv2

import (
	// "context"
	// "net/http"
	"log"
	// "time"

	clientv2 "go.etcd.io/etcd/client/v2"
)

func Connect(endpoints []string) (clientv2.Client, error) {

	cli, err := clientv2.New(clientv2.Config{
		Endpoints: endpoints,
		Transport: clientv2.DefaultTransport,
	})

	if err != nil {
		log.Fatal(err)
	}

	return cli, err
}
