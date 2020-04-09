package router

import (
	"net/url"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/log"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

func authorize(c *gin.Context) {
	params := c.Request.URL.Query().Encode()
	// cookie中未找到session，跳转到登录页面
	session, _ := c.Cookie("session")
	log.Info("session:", session)
	if session == "" {
		loginUrl, _ := url.Parse("http://localhost:2233/account/static/login")
		web.Redirect(c, loginUrl.String()+"?"+params)
		return
	}
	// 已登录，获取用户id
	userId := checkAndGetUserId(session)
	if userId <= 0 {
		loginUrl, _ := url.Parse("http://localhost:2233/account/static/login")
		web.Redirect(c, loginUrl.String()+"?"+params)
		return
	}

	req := new(model.AuthorizeReq)
	err := c.ShouldBind(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	locationUrl, err := srv.Authorize(c, userId, req.ClientKey, req.ResponseType, req.RedirectUri, req.Scope, req.State)
	if err != nil {
		web.JSON(c, nil, err)
		return
	}
	log.Info("locationUrl:", locationUrl)
	web.Redirect(c, locationUrl)
}
