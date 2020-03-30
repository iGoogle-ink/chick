package router

import (
	"fmt"

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
	//srv.Authorize(req.ClientKey, req.ResponseType, req.RedirectUri, req.State)
	c.Redirect(302, "http://localhost:8082/account/oauth/callback?code=nihaowoshicode&state=qwer1234asdf")
}

func callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println("code:", code)
	fmt.Println("state:", state)
}
