// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package postprocessor

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/ravin66/packer-plugin-etcd/internal/etcdv2"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Endpoint string `mapstructure:"endpoint"`
	Key      string `mapstructure:"key"`
	Value    string `mapstructure:"value"`
	Method   string `mapstructure:"method"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec { return p.config.FlatMapstructure().HCL2Spec() }

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.post-processor.etcd",
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
	if p.config.Username != "" {
		cfg.UseAuth = true
		cfg.Password = p.config.Password
		cfg.Username = p.config.Username
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
			return source, false, false, err
		}

		ui.Message("Successfully set key: " + p.config.Key + " to value: " + p.config.Value)
		ui.Message("Key: " + p.config.Key)
		ui.Message("Value: " + p.config.Value)

	case "get":
		val, err := etcdv2.Get(eApi, p.config.Key)
		if err != nil {
			ui.Error("Failed to get value: " + err.Error())
			return source, false, false, err
		}
		ui.Message("Value retrieved: " + val)

	case "delete":
		err := etcdv2.Del(eApi, p.config.Key)
		if err != nil {
			ui.Error("Failed to delete key: " + err.Error())
			return source, false, false, err
		}
		ui.Message("Successfully deleted key: " + p.config.Key)

	default:
		err := fmt.Errorf("unsupported method: %s", p.config.Method)
		ui.Error(err.Error())
		return source, false, false, err
	}

	return source, true, true, nil
}
