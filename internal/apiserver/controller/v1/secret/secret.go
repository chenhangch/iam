package secret

import (
	srvv1 "github.com/chenhangch/iam/internal/apiserver/service/v1"
	"github.com/chenhangch/iam/internal/apiserver/store"
)

// SecretController 创建一个密钥handler
type SecretController struct {
	srv srvv1.Service
}

// NewSecretController creates a secret handler.
func NewSecretController(factory store.Factory) *SecretController {
	return &SecretController{
		srv: srvv1.NewService(factory),
	}
}
