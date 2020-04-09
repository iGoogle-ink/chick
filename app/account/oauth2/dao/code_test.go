package dao

import (
	"context"
	"testing"

	"chick/pkg/log"
	"github.com/smartystreets/goconvey/convey"
)

func TestDao_GetCodeInfoByCode(t *testing.T) {
	convey.Convey("GetCodeInfoByCode", t, func(c convey.C) {
		var (
			ctx  = context.Background()
			code = "06b19edc66d9bf01851c714d7f26243a"
		)
		codeInfo, err := d.CacheAuthCode(ctx, code)
		c.Convey("GetCodeInfoByCode error should be nil", func(c convey.C) {
			c.So(err, convey.ShouldBeNil)
		})
		log.Info(codeInfo)
	})
}
