package router

import (
	"github.com/gin-gonic/gin"
	system "github.com/titrxw/smart-home-server/router/System"
)

type Router struct {
	system.RouteGroup
}

func (this *Router) Register(router *gin.Engine) {
	router.Static("/static", "./public/static")

	this.RegisterBaseRouteGroup(router)
}

var GRouter = new(Router)
