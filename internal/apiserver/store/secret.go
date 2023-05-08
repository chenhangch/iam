package store

import (
	"context"

	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type SecretStore interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
}
