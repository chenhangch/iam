package apiserver

import "github.com/chang144/iam/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrePareRun().Run()
}
