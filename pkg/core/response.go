package core

import (
	"net/http"

	"github.com/chenhangch/golunzi/errors"
	"github.com/gin-gonic/gin"
	
	log "golang.org/x/exp/slog"
)

// ErrResponse 定义了错误返回格式
type ErrResponse struct {
	Code int `json:"backend"`

	Message string `json:"message"`

	Reference string `json:"reference,omitempty"`
}

func WriteResponse(c *gin.Context, err error, date interface{}) {
	if err != nil {
		log.Error("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})
		return
	}
	c.JSON(http.StatusOK, date)
}
