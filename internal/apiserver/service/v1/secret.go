package v1

import (
	"context"
	"github.com/chang144/golunzi/errors"
	"github.com/chang144/iam/internal/apiserver/store"
	"github.com/chang144/iam/internal/pkg/code"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type SecretSrv interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
}

type secretService struct {
	store store.Factory
}

var _ SecretSrv = (*secretService)(nil)

func newSecrets(srv *service) *secretService {
	return &secretService{
		store: srv.store,
	}
}

// Create 创建密钥
func (s secretService) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error {
	if err := s.store.Secrets().Create(ctx, secret, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}
