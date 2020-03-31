package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDao_CloudUserInfo(t *testing.T) {
	convey.Convey("CloudUserInfo", t, func(c convey.C) {
		var (
			ctx  = context.Background()
			name = "kx"
		)
		user, err := d.CloudUserInfo(ctx, name)
		c.Convey("CloudUserInfo err should be nil", func(c convey.C) {
			c.So(err, convey.ShouldBeNil)
			fmt.Println("user:", user)
		})
	})
}
