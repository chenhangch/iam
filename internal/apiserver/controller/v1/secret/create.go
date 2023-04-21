package secret

import (
	"github.com/chang144/ciam/internal/pkg/code"
	"github.com/chang144/ciam/pkg/core"
	"github.com/chang144/ciam/pkg/log"
	"github.com/chang144/golunzi/errors"
	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
)

const maxSecretCout = 10

func (s *SecretController) Create(c *gin.Context) {
	log.L(c).Info("Creating secret")

	var r v1.Secret
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

}
