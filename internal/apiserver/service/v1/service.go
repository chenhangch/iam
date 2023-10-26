package v1

import "github.com/chang144/iam/internal/apiserver/store"

type Service interface {
	Secrets() SecretSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{store: store}
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}
