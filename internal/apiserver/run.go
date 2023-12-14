package apiserver

import "github.com/chenhangch/iam/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrePareRun().Run()
}
