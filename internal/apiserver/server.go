package apiserver

import (
	"github.com/chenhangch/golunzi/shutdown"
	"github.com/chenhangch/golunzi/shutdown/posixsignal"
	"github.com/chenhangch/iam/internal/apiserver/config"
	"github.com/chenhangch/iam/internal/apiserver/store/mysql"
	"github.com/chenhangch/iam/internal/pkg/server"
	"github.com/chenhangch/iam/pkg/log"

	genericapiserver "github.com/chenhangch/iam/internal/pkg/server"
)

type apiServer struct {
	gs *shutdown.GracefulShutdown

	// TODO : Redis && grpc
	genericAPIServer *server.GenericAPIServer
}

// preparedAPIServer 表示api Server准备完成
type preparedAPIServer struct {
	*apiServer
}

// ExtraConfig defines extra configuration
type ExtraConfig struct {
	Addr       string
	MaxMsgSize int
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	return &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}, nil
}

// PrePareRun prePare to run
// init Redis and Add close()
func (s *apiServer) PrePareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	// TODO:initRedis()

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}
		// TODO: 关闭gRPC
		s.genericAPIServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	// TODO: go s.gRPCAPIServer.Run()

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}

// buildGenericConfig creates generic config [config -> logicServer]
func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewNilConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	// TODO: FeatureOptions
	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	return
}
