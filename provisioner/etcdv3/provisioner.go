package provisioner

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/ravin66/packer-plugin-etcd/internal/etcdv3"
)

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
	Key      string `mapstructure:"key"`
	Value    string `mapstructure:"value"`
	Method   string `mapstructure:"method"`

	ctx interpolate.Context
}

type Provisioner struct {
	config Config
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.provisioner.scaffolding",
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator, generatedData map[string]interface{}) error {
	// ui.Say(fmt.Sprintf("provisioner mock: %s", p.config.MockOption))
	// return nil
	cli, err := etcdv3.Connect([]string{p.config.Endpoint}, 5*time.Second)
	if err != nil {
		ui.Error("Failed to connect to etcd: " + err.Error())
		return err
	}
	defer cli.Close()

	switch strings.ToLower(p.config.Method) {
	case "put":
		err := etcdv3.Put(cli, p.config.Key, p.config.Value)
		if err != nil {
			ui.Error("Failed to PUT value: " + err.Error())
			return err
		}
		ui.Message("Successfully set key: " + p.config.Key + " to value: " + p.config.Value)

	case "get":
		val, err := etcdv3.Get(cli, p.config.Key)
		if err != nil {
			ui.Error("Failed to GET value: " + err.Error())
			return err
		}
		ui.Message("Value retrieved: " + val)

	case "delete":
		err := etcdv3.Del(cli, p.config.Key)
		if err != nil {
			ui.Error("Failed to DELETE key: " + err.Error())
			return err
		}
		ui.Message("Successfully deleted key: " + p.config.Key)

	default:
		err := fmt.Errorf("unsupported method: %s", p.config.Method)
		ui.Error(err.Error())
		return err
	}

	return nil

}
