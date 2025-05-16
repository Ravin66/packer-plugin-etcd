packer {
  required_plugins {
    etcdv3 = {
      version = ">0.1.1"
      source  = "github.com/ravin66/etcdv3"
    }
  }
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "etcdv3-etcd" {
    endpoint = ""
    key      = ""
    value    = ""
    method   = ""
  }

  post-processor "etcdv3-etcd" {
    endpoint = ""
    key      = ""
    value    = ""
    method   = ""
  }
}
