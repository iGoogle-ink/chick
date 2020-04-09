package web

import (
	"chick/pkg/log"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	Gin  *gin.Engine
	port string
}

func InitServer(port string) *Engine {
	engine := &Engine{
		Gin:  gin.Default(),
		port: port,
	}
	return engine
}

func (e *Engine) Start() {
	go func() {
		if err := e.Gin.Run(e.port); err != nil {
			log.Panicf("web server port(%s) run error(%+v).", e.port, err)
		}
	}()
}
