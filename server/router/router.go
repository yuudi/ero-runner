package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuudi/ero-runner/server/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	return r
}
