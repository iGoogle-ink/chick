package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
)

var mKey = map[string]bool{
	"sd": false,
}

func Verify(c *gin.Context) {
	req := c.Request
	h := req.Header

	sappkey := h.Get("appkey")
	if _, ok := mKey[sappkey]; ok { // only allowed keys can access
		return
	}
	log.Warn("IllegalKey! %s", sappkey)
}
