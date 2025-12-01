package etcdv2

import (
	"context"
	"log"
	"time"

	clientv2 "go.etcd.io/etcd/client/v2"
)

type EtcdOptions struct {
	Endpoints []string
	Username  string
	Password  string
	UseAuth   bool
}

func Connect(config EtcdOptions) (clientv2.KeysAPI, error) {

	cfg := clientv2.Config{
		Endpoints:               config.Endpoints,
		Transport:               clientv2.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	}

	if config.UseAuth && (config.Username != "" || config.Password != "") {
		cfg.Username = config.Username
		cfg.Password = config.Password
	} else if config.UseAuth && (config.Username == "" || config.Password == "") {
		log.Fatal("Use Auth has been set to true but no username or password provided.")
	}

	cli, err := clientv2.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	return clientv2.NewKeysAPI(cli), nil
}

func Put(api clientv2.KeysAPI, key, value string) error {
	_, err := api.Set(context.Background(), key, value, nil)

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
