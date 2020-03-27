package router

import (
	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

func authorize(c *gin.Context) {
	req := new(model.AuthorizeReq)
	err := c.ShouldBindQuery(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	srv.Authorize(req.ClientKey, req.ResponseType, req.RedirectUri, req.State)
}
