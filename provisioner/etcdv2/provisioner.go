package provisioner

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/ravin66/packer-plugin-etcd/internal/etcdv2"
)

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
	Key      string `mapstructure:"key"`
	Value    string `mapstructure:"value"`
	Method   string `mapstructure:"method"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`

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
		PluginType:         "packer.provisioner.etcd",
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

	cfg := etcdv2.EtcdOptions{
		Endpoints: []string{p.config.Endpoint},
		UseAuth:   false,
	}

	// Check to see if environment variables are set and override
	if os.Getenv("ETCD_USERNAME") != "" {
		cfg.UseAuth = true
		cfg.Username = os.Getenv("ETCD_USERNAME")
	}

	if os.Getenv("ETCD_PASSWORD") != "" {
		cfg.Password = os.Getenv("ETCD_PASSWORD")
	}

	// Check if the username is configured in packer.  This will override the environment variable.
	// check if the username is configured in packer.  This will override the environment variable
	if p.config.Username != "" {
		cfg.UseAuth = true
		cfg.Username = p.config.Username
	}

	if p.config.Password != "" {
		cfg.UseAuth = true
		cfg.Password = p.config.Password
	}

	if cfg.UseAuth && (p.config.Password == "" || p.config.Username == "") {
		ui.Error("Authentication is set to True, but there is no username or password provided.")
		return nil
	}

	// Connect to etcdv2
	eApi, err := etcdv2.Connect(cfg)

	if err != nil {
		ui.Error("Failed to connect to etcd: " + err.Error())
	}

	switch strings.ToLower(p.config.Method) {
	case "put":
		err := etcdv2.Put(eApi, p.config.Key, p.config.Value)

		if err != nil {
			ui.Error("Failed to set key: " + err.Error())
			return err
		}

		ui.Message("Successfully set key: " + p.config.Key + " to value: " + p.config.Value)
		ui.Message("Key: " + p.config.Key)
		ui.Message("Value: " + p.config.Value)

	case "get":
		val, err := etcdv2.Get(eApi, p.config.Key)
		if err != nil {
			ui.Error("Failed to get value: " + err.Error())
			return err
		}
		ui.Message("Value retrieved: " + val)

	case "delete":
		err := etcdv2.Del(eApi, p.config.Key)
		if err != nil {
			ui.Error("Failed to delete key: " + err.Error())
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
