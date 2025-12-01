package main

import (
	"fmt"
	"os"

	datav3 "github.com/ravin66/packer-plugin-etcd/datasource/etcdv3"
	ppdv3 "github.com/ravin66/packer-plugin-etcd/post-processor/etcdv3"
	provv3 "github.com/ravin66/packer-plugin-etcd/provisioner/etcdv3"

	datav2 "github.com/ravin66/packer-plugin-etcd/datasource/etcdv2"
	ppdv2 "github.com/ravin66/packer-plugin-etcd/post-processor/etcdv2"
	provv2 "github.com/ravin66/packer-plugin-etcd/provisioner/etcdv2"

	version "github.com/ravin66/packer-plugin-etcd/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()

	pps.RegisterDatasource("etcdv3", new(datav3.Datasource))
	pps.RegisterPostProcessor("etcdv3", new(ppdv3.PostProcessor))
	pps.RegisterProvisioner("etcdv3", new(provv3.Provisioner))

	pps.RegisterDatasource("etcdv2", new(datav2.Datasource))
	pps.RegisterPostProcessor("etcdv2", new(ppdv2.PostProcessor))
	pps.RegisterProvisioner("etcdv2", new(provv2.Provisioner))

	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
