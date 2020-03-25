package http

import (
	"net/http"

	"chick/errno"
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, data interface{}, err error) {
	e := errno.AnalyseError(err)

	rsp := &struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Code:    e.Code(),
		Message: e.Message(),
		Data:    data,
	}

	c.JSON(http.StatusOK, rsp)
}
