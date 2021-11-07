package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/titrxw/smart-home-server/app/Controller"
)

func Register(router *gin.Engine) {
	router.GET("/", new(controller.IndexController).Index)
	router.Static("/static", "./public/static")
}
