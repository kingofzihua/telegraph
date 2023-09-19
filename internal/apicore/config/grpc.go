package config

import (
	"github.com/spf13/pflag"
)

type GrpcConfig struct {
	Address string `json:"address"`
}

// NewServerConfig 初始化 config
func NewGrpcConfig() *GrpcConfig {
	return &GrpcConfig{
		Address: ":8001",
	}
}

// Validate 验证数据
func (s *GrpcConfig) Validate() []error {
	errors := []error{}

	return errors
}

// AddFlags 绑定 pflag
func (s *GrpcConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Address, "grpc.address", s.Address, "on http listen address")
}
