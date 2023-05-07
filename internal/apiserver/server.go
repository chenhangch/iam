package apiserver

import (
	"github.com/chang144/ciam/internal/apiserver/config"
	"github.com/chang144/golunzi/shutdown"
	"github.com/chang144/golunzi/shutdown/posixsignal"
)

type apiServer struct {
	gs *shutdown.GracefulShutdown
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	return &apiServer{gs: gs}, nil
}
