package config

import (
	"github.com/go-ostrich/pkg/server"
	"github.com/spf13/pflag"
)

type ServerConfig struct {
	Mode    string `json:"mode"`
	Address string `json:"address"`
}

// NewServerConfig 初始化 config
func NewServerConfig() *ServerConfig {
	// 默认的配置
	def := server.NewConfig()
	return &ServerConfig{
		Address: def.Address,
	}
}

// ApplyTo 将从命令行或者配置文件中的配置应用到 服务支持的配置
func (s *ServerConfig) ApplyTo(c *server.Config) error {
	// 如果当前的配置不是空的，就覆盖
	if s.Address != "" {
		c.Address = s.Address
	}
	if s.Mode != "" {
		c.Mode = s.Mode
	}
	return nil
}

// Validate 验证数据
func (s *ServerConfig) Validate() []error {
	errors := []error{}

	return errors
}

// AddFlags 绑定 pflag
func (s *ServerConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Address, "server.address", s.Address, "on http listen address")
}
