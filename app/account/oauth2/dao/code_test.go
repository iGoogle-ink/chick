package dao

import (
	"context"
	"fmt"
	"testing"
	"time"

	"chick/app/account/oauth2/model"
	xtime "chick/pkg/time"
	"github.com/google/uuid"
	"github.com/smartystreets/goconvey/convey"
)

func TestDao_GetCodeInfoByCode(t *testing.T) {
	convey.Convey("GetCodeInfoByCode", t, func(c convey.C) {
		var (
			ctx  = context.Background()
			code = "06b19edc66d9bf01851c714d7f26243a"
		)
		codeInfo, err := d.GetCodeInfoByCode(ctx, code)
		c.Convey("GetCodeInfoByCode error should be nil", func(c convey.C) {
			c.So(err, convey.ShouldBeNil)
			fmt.Println("codeInfo:", codeInfo)
		})
	})
}

func TestDao_InsertCode(t *testing.T) {
	convey.Convey("InsertCode", t, func(c convey.C) {
		var (
			ctx  = context.Background()
			code = &model.OauthAuthCode{
				ClientId:    1,
				UserId:      1,
				Code:        uuid.New().String(),
				RedirectUri: "http://www",
				ExpiresAt:   xtime.Time(time.Now().Add(time.Minute * 5).Unix()),
				Scope:       "all",
			}
		)
		err := d.InsertCode(ctx, code)
		c.Convey("InsertCode error should be nil", func(c convey.C) {
			c.So(err, convey.ShouldBeNil)
		})
	})
}
