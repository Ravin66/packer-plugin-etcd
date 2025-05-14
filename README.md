# Packer Plugin ETCD

This Packer plugin is used to interact with ETCDv3 and ETCDv2 (WIP).

# Usage
Configure the packer template to require the required release version of this plugin. 

For example:

```hcl
packer {
  required_plugins {
    etcdv3 = {
      version = "0.1.1"
      source  = "github.com/ravin66/etcdv3"
    }
  }
}
```

Initialise the packer template using `init` to install the plug in

```
packer init <name of template>.pkr.hcl
```

Currently, this plugin has a provisioner and post-processor which can be used.  At this stage the plugin doesn't do auth but this will be fixed shortly.

Create keys:
```hcl
provisioner "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/Provisioner"
    value    = "This is a provsioner key"
    method   = "put"
}

  post-processor "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/post-process"
    value    = "This is a post process key"
    method   = "put"
  }
```

Get keys:
```hcl
provisioner "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/Provisioner"
    value    = "This is a provsioner key"
    method   = "get"
}

  post-processor "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/post-process"
    value    = "This is a post process key"
    method   = "get"
  }
```

Delete keys: 
```hcl
provisioner "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/Provisioner"
    value    = "This is a provsioner key"
    method   = "delete"
}

  post-processor "etcdv3-etcd" {
    endpoint = "localhost:2379"
    key      = "/test/post-process"
    value    = "This is a post process key"
    method   = "delete"
  }
```

Current this is all the plugin can do.  There are plans to expand on the above.

# Development

TBC