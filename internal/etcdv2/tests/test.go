package main

import (
	"fmt"
	"log"

	clientv2 "github.com/ravin66/packer-plugin-etcd/internal/etcdv2"
)

func main() {
	cli, err := clientv2.Connect([]string{"http://192.168.0.13:2379"})
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	kapi := clientv2.NewKeysAPI(cli)

	key := "/example/key"
	value := "hello etcdv2"

	if err := setKey(kapi, key, value); err != nil {
		log.Fatalf("Set error: %v", err)
	}

	val, err := getKey(kapi, key)
	if err != nil {
		log.Fatalf("Get error: %v", err)
	}
	fmt.Printf("Value retrieved: %s\n", val)

	if err := deleteKey(kapi, key); err != nil {
		log.Fatalf("Delete error: %v", err)
	}
}
