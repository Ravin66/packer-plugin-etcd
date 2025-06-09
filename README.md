# packer-plugin-etcd

A [Packer](https://www.packer.io/) plugin for reading from and writing to an already configured [etcd](https://etcd.io/) cluster or single instance. This plugin enables Packer workflows to interact with etcd for retrieving or storing data during image builds. Useful for dynamic configuration, distributed state sharing, or secrets management.

> **Note:** This plugin does **not** provision or configure etcd clusters. It assumes you already have an accessible etcd instance.

## Usage

Configure the Packer template to require the desired release version of this plugin. 

For example:

```hcl
packer {
  required_plugins {
    etcd = {
      version = ">= 0.1.1"
      source  = "github.com/ravin66/etcdv3"
    }
  }
}
```

Initialize the Packer template using `init` to install the plugin:

```sh
packer init <name of template>.pkr.hcl
```

Alternatively, you can install the plugin directly using a local path:

```sh
packer plugins install --path "<path>\Packer-plugin-etcd_v0.0.2_x5.0_linux_amd64" "github.com/ravin66/etcd"
```

## Authentication

Both etcd v3 and v2 support authentication with Packer. This can be configured via environment variables or in the HCL settings.

> **Note:** If you configure both the username and password in environment variables and in HCL, the HCL configuration will take precedence.

The following environment variables can be set in your build environment:

- `ETCD_USERNAME`: etcd username
- `ETCD_PASSWORD`: etcd password

For example, on Linux/macOS:

```sh
export ETCD_USERNAME=your-username
export ETCD_PASSWORD=your-password
```

Or on Windows:

```sh
set ETCD_USERNAME=your-username
set ETCD_PASSWORD=your-password
```

**HCL Example:**
```hcl
provisioner "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  value    = "127.0.0.1"
  username = "your-username"
  password = "your-password"
  method   = "put"
}
```

## ETCD API V3

The etcd v3 client currently supports provisioner and post-processor components. Expansion to other areas is planned.

### Create keys:
```hcl
provisioner "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  value    = "127.0.0.1"
  method   = "put"
}

post-processor "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  value    = "127.0.0.1"
  method   = "put"
}
```

### Get keys:
```hcl
provisioner "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "get"
}

post-processor "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "get"
}
```

### Delete keys: 
```hcl
provisioner "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "delete"
}

post-processor "etcd-etcdv3" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "delete"
}
```

## ETCD API V2

Currently, the etcd v2 client only has the post-processor configured. Expansion to other areas is planned.

### Create keys:
```hcl
post-processor "etcd-etcdv2" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  value    = "127.0.0.1"
  method   = "put"
}
```

### Get keys:
```hcl
post-processor "etcd-etcdv2" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "get"
}
```

### Delete keys:
```hcl
post-processor "etcd-etcdv2" {
  endpoint = "http://localhost:2379"
  key      = "post-processor"
  method   = "delete"
}
```

Currently, this is all the plugin can do. There are plans to expand on the above.

## Prerequisites

- An already running and accessible etcd cluster or single instance.
- No additional dependencies required other than etcd being reachable.

## License

This project is open source and released under the [MIT License](LICENSE).

## Development

*Contributions are welcome!*