package postprocessor

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Endpoint string `mapstructure:"endpoint"`
	Key      string `mapstructure:"key"`
	Value    string `mapstructure:"value"`
	Method   string `mapstructure:"method"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec { return p.config.FlatMapstructure().HCL2Spec() }

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.post-processor.scaffolding",
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

func (p *PostProcessor) PostProcess(ctx context.Context, ui packersdk.Ui, source packersdk.Artifact) (packersdk.Artifact, bool, bool, error) {
	// cli, err := etcdv2.Connect([]string{p.config.Endpoint}, 5*time.Second)
	// if err != nil {
	// 	ui.Error("Failed to connect to etcd: " + err.Error())
	// 	return source, false, false, err
	// }
	// defer cli.Close()

	// switch strings.ToLower(p.config.Method) {
	// case "put":
	// 	err := etcdv2.Put(cli, p.config.Key, p.config.Value)
	// 	if err != nil {
	// 		ui.Error("Failed to set key: " + err.Error())
	// 		return source, false, false, err
	// 	}
	// 	ui.Message("Successfully set key: " + p.config.Key + " to value: " + p.config.Value)
	// 	ui.Message("Key: " + p.config.Key)
	// 	ui.Message("Value: " + p.config.Value)

	// case "get":
	// 	val, err := etcdv2.Get(cli, p.config.Key)
	// 	if err != nil {
	// 		ui.Error("Failed to get value: " + err.Error())
	// 		return source, false, false, err
	// 	}
	// 	ui.Message("Value retrieved: " + val)

	// case "del":
	// 	err := etcdv2.Del(cli, p.config.Key)
	// 	if err != nil {
	// 		ui.Error("Failed to delete key: " + err.Error())
	// 		return source, false, false, err
	// 	}
	// 	ui.Message("Successfully deleted key: " + p.config.Key)

	// default:
	// 	err := fmt.Errorf("unsupported method: %s", p.config.Method)
	// 	ui.Error(err.Error())
	// 	return source, false, false, err
	// }
	ui.Say(fmt.Sprintf("post-processor mock: %s", p.config.Endpoint))
	return source, true, true, nil
}
