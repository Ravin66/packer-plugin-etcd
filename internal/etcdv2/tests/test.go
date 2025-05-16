package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ravin66/packer-plugin-etcd/internal/etcdv2"
)

func main() {
	inputKeys := flag.String("keys", "", "Comma-separtated list of keys.")
	etcdhost := flag.String("etcdHost", "", "Etcd Hostname to get keys.")
	flag.Parse()

	cfg := etcdv2.EtcdOptions{
		Endpoints: []string{*etcdhost},
	}

	eApi, err := etcdv2.Connect(cfg)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	keys := strings.Split(*inputKeys, ",")

	for _, key := range keys {
		// fmt.Printf(key)
		value, err := etcdv2.Get(eApi, key)

		if err != nil {
			fmt.Printf("Failed to get key: %v", err)
		} else {
			fmt.Printf("Key: %s - Value %s\n", key, value)
		}
	}

}
