package apiserver

import (
	"fmt"

	"github.com/go-ostrich/pkg/server"
	"github.com/go-ostrich/pkg/shutdown"
	"github.com/go-ostrich/pkg/shutdown/shutdownmanagers/posixsignal"

	"github.com/kingofzihua/telegraph/internal/apiserver/config"
)

type apiServer struct {
	gs     *shutdown.GracefulShutdown
	server *server.GinServer
}

// PrepareRun .
func (s *apiServer) PrepareRun() preparedAPIServer {
	// init router
	err := installRouters(s.server.Engine)
	if err != nil {
		panic(fmt.Errorf("install router error: %+v", err))
	}
	return preparedAPIServer{s}
}

// wrap apiServer
type preparedAPIServer struct {
	*apiServer
}

// Run server
func (s preparedAPIServer) Run() error {
	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		return err
	}
	// start server
	return s.server.Run()
}

// create api server
func createAPIServer(cfg *config.Config) (*apiServer, error) {
	// 优雅停止
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	srv, err := buildServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("apiserver build config error: %+v", err)
	}

	s := &apiServer{
		gs:     gs,
		server: srv,
	}
	return s, nil
}

// from config build server
func buildServer(cfg *config.Config) (*server.GinServer, error) {
	// 默认的配置文件
	serverConfig := server.NewConfig()

	// 将自定义的配置参数应用
	if err := cfg.Server.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	return serverConfig.Complete().NewGinServer(), nil
}
