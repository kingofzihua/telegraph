package server

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Mode    string `json:"mode"`
	Address string `json:"address"`
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Mode:    gin.ReleaseMode,
		Address: "127.0.0.1:8080",
	}
}

// CompletedConfig is the completed configuration
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

func (c CompletedConfig) NewGinServer() *GinServer {
	gin.SetMode(c.Mode)

	return &GinServer{
		Engine:  gin.New(),
		Address: c.Address,
	}
}
