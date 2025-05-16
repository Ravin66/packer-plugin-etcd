package main

import (
	"fmt"
	"os"

	ppdv3 "github.com/ravin66/packer-plugin-etcd/post-processor/etcdv3"
	provv3 "github.com/ravin66/packer-plugin-etcd/provisioner/etcdv3"

	ppdv2 "github.com/ravin66/packer-plugin-etcd/post-processor/etcdv2"
	// provv2 "github.com/ravin66/packer-plugin-etcd/provisioner/etcdv2"

	version "github.com/ravin66/packer-plugin-etcd/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	// pps.RegisterBuilder("my-builder", new(scaffolding.Builder))
	pps.RegisterPostProcessor("etcdv3", new(ppdv3.PostProcessor))
	pps.RegisterProvisioner("etcdv3", new(provv3.Provisioner))
	pps.RegisterPostProcessor("etcdv2", new(ppdv2.PostProcessor))
	// pps.RegisterProvisioner("etcdv2", new(provv2.Provisioner))

	// pps.RegisterDatasource("my-datasource", new(scaffoldingData.Datasource))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
