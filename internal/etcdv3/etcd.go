package etcdv3

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdOptions struct {
	Endpoints []string
	Username  string
	Password  string
	UseAuth   bool
}

func Connect(config EtcdOptions) (*clientv3.Client, error) {

	cfg := clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: 5 * time.Second,
	}

	if config.UseAuth && (config.Username != "" || config.Password != "") {
		cfg.Username = config.Username
		cfg.Password = config.Password
	} else if config.UseAuth && (config.Username == "" || config.Password == "") {
		log.Fatal("Use Auth has been set to true but no username or password provided.")
	}

	cli, err := clientv3.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cli, err
}

func Disconnect(cli *clientv3.Client) {
	err := cli.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func Put(cli *clientv3.Client, key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cli.Put(ctx, key, value)
	return err
}

func Get(cli *clientv3.Client, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := cli.Get(ctx, key)
	if err != nil {
		return "", err
	}

	if len(resp.Kvs) > 0 {
		return string(resp.Kvs[0].Value), nil
	}

	return "", nil
}

func Del(cli *clientv3.Client, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cli.Delete(ctx, key)
	return err
}
