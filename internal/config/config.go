// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	cliflag "github.com/go-ostrich/pkg/cli/flag"

	"github.com/go-ostrich/pkg/json"
)

// Config is the running configuration structure of the IAM pump service.
type Config struct {
	Server *ServerConfig
}

// NewOptions app opts
func NewOptions() *Config {
	return &Config{}
}

// Validate 验证配置合法性
func (c *Config) Validate() []error {
	var errs []error

	return errs
}

// Flags returns flags for a specific APIServer by section name.
func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	return fss
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)

	return string(data)
}

// Complete set default Options.
func (c *Config) Complete() error {
	return nil
}
