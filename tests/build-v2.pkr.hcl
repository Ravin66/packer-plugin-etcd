packer {
  required_plugins {
    etcd = {
      version = ">=0.1.2"
      source  = "github.com/ravin66/etcd"
    }
  }
}

locals {
  etcdhost = "http://localhost:2379"
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  post-processor "etcd-etcdv2" {
    endpoint = local.etcdhost
    key      = "post-processor"
    value    = "127.0.0.1"
    method   = "put"
  }

  post-processor "etcd-etcdv2" {
    endpoint = local.etcdhost
    key      = "post-processor"
    method   = "get"
  }

  post-processor "etcd-etcdv2" {
    endpoint = local.etcdhost
    key      = "post-processor"
    method   = "delete"
  }
}
