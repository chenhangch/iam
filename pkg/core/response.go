package core

import (
	"github.com/chang144/golunzi/errors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
)

// ErrResponse 定义了错误返回格式
type ErrResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`

	Reference string `json:"reference,omitempty"`
}

func WriteResponse(c *gin.Context, err error, date interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
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
