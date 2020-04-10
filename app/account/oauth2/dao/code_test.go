package dao

import (
	"context"
	"testing"

	"chick/app/account/oauth2/model"
	"chick/pkg/log"

	"github.com/smartystreets/goconvey/convey"
)

func TestDao_AddCacheAuthCode(t *testing.T) {
	convey.Convey("AddCacheAuthCode", t, func() {
		var (
			ctx   = context.Background()
			code  = "06b19edc66d9bf01851c714d7f26243a"
			codes = &model.CacheAuthCode{
				ClientKey:   "123123",
				UserId:      1,
				Code:        code,
				RedirectUri: "Http",
				Expires:     36000,
				Scope:       "all",
			}
		)
		err := d.AddCacheAuthCode(ctx, codes)
		convey.So(err, convey.ShouldBeNil)
	})
}

func TestDao_GetCodeInfoByCode(t *testing.T) {
	convey.Convey("GetCodeInfoByCode", t, func(c convey.C) {
		var (
			ctx  = context.Background()
			code = "06b19edc66d9bf01851c714d7f26243a"
		)
		codeInfo, err := d.CacheAuthCode(ctx, code)
		c.So(err, convey.ShouldBeNil)
		log.Info(codeInfo)
	})
}
