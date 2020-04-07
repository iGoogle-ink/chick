package server

import (
	"context"
	"fmt"
	"time"

	"chick/pkg/util"
	"github.com/micro/go-micro/v2/server"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%s]: server [%s] received request: %s", time.Now().Format(util.TimeLayout), req.Service(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}
