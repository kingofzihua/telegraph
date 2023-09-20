package apicore

import (
	"net"

	"github.com/go-ostrich/pkg/log"

	"github.com/kingofzihua/telegraph/internal/apicore/service"

	"github.com/go-ostrich/pkg/shutdown"
	"github.com/go-ostrich/pkg/shutdown/shutdownmanagers/posixsignal"
	"google.golang.org/grpc"

	"github.com/kingofzihua/telegraph/internal/apicore/config"
	v1 "github.com/kingofzihua/telegraph/proto/server/v1"
)

type grpcServer struct {
	gs          *shutdown.GracefulShutdown
	server      *grpc.Server
	lis         net.Listener
	mysqlConfig *config.MySQLConfig
}

// PrepareRun Pre Run
func (s *grpcServer) PrepareRun() *preparedGrpcServer {
	logger := log.Default()

	psrv := service.NewPageService(logger)

	// register server
	v1.RegisterPageServiceServer(s.server, psrv)

	return &preparedGrpcServer{s}
}

// wrap apiServer
type preparedGrpcServer struct {
	*grpcServer
}

// Run grpc server.
func (s preparedGrpcServer) Run() error {
	// start grpc server
	return s.server.Serve(s.lis)
}

func createGrpcServer(cfg *config.Config) (*grpcServer, error) {
	// 优雅停止
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	// 创建 GRPC Server 实例
	grpcsrv := grpc.NewServer()

	// listener
	lis, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		return nil, err
	}

	s := &grpcServer{
		gs:     gs,
		server: grpcsrv,
		lis:    lis,
	}
	return s, nil
}
