// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package etcd

import (
	"os"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"

	"github.com/ravin66/packer-plugin-etcd/internal/etcdv2"
)

type EtcdOptions struct {
	Endpoints []string
	Username  string
	Password  string
	UseAuth   bool
}

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
	Key      string `mapstructure:"key"`
	Value    string `mapstructure:"value"`
	Method   string `mapstructure:"method"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	Key   string `mapstructure:"key"`
	Value string `mapstructure:"value"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	empytOutput := hcl2helper.HCL2ValueFromConfig(DatasourceOutput{}, d.OutputSpec())

	cfg := etcdv2.EtcdOptions{
		Endpoints: []string{d.config.Endpoint},
		UseAuth:   false,
	}

	// Check to see if environment variables are set and override.
	if os.Getenv("ETCD_USERNAME") != "" {
		cfg.UseAuth = true
		cfg.Username = os.Getenv("ETCD_USERNAME")
	}

	if os.Getenv("ETCD_PASSWORD") != "" {
		cfg.UseAuth = true
		cfg.Password = os.Getenv("ETCD_PASSWORD")
	}

	// check if the username is configured in packer.  This will override the environment variable
	if d.config.Username != "" {
		cfg.UseAuth = true
		cfg.Username = d.config.Username
	}

	if d.config.Password != "" {
		cfg.UseAuth = true
		cfg.Password = d.config.Password
	}

	if cfg.UseAuth && (d.config.Password == "" || d.config.Username == "") {
		return empytOutput, nil
	}

	eApi, err := etcdv2.Connect(cfg)

	if err != nil {
		return empytOutput, err
	}

	val, err := etcdv2.Get(eApi, d.config.Key)

	if err != nil {
		return empytOutput, err
	}

	output := DatasourceOutput{
		Key:   d.config.Key,
		Value: val,
	}
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
