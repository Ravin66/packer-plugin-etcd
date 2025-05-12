package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	clientv3 "github.com/ravin66/packer-plugin-etcd/internal/etcdv3"
)

func main() {

	inputKeys := flag.String("keys", "", "Comma-separtated list of keys.")
	etcdhost := flag.String("etcdHost", "", "Etcd Hostname to get keys.")
	flag.Parse()

	keys := strings.Split(*inputKeys, ",")

	cli, err := clientv3.Connect([]string{*etcdhost}, 5*time.Second)

	if err != nil {
		fmt.Errorf("Failed to connect to etcd: %v", err)
	}

	defer cli.Close()

	for _, key := range keys {
		// fmt.Printf(key)
		value, err := clientv3.Get(cli, key)

		if err != nil {
			fmt.Printf("Failed to get key: %v", err)
		} else {
			fmt.Printf("Key: %s - Value %s\n", key, value)
		}
	}
}
