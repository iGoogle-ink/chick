package router

import (
	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

func token(c *gin.Context) {
	req := new(model.AccessTokenReq)
	err := c.ShouldBind(&req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	token, err := srv.GetAccessToken(c, req)
	if err != nil {
		web.JSON(c, nil, err)
		return
	}
	web.JSON(c, token, nil)
}

func refreshToken(c *gin.Context) {
	req := model.OauthRefreshTokenReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	token, err := srv.GetNewToken(c, &req)
	if err != nil {
		web.JSON(c, nil, err)
		return
	}
	web.JSON(c, token, nil)
}
