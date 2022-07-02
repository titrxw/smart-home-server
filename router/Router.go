package router

import (
	"github.com/gin-gonic/gin"
	frontend "github.com/titrxw/smart-home-server/router/Frontend"
	system "github.com/titrxw/smart-home-server/router/System"
)

type Router struct {
	system.SysRouteGroup
	frontend.FrontendRouteGroup
}

func (r *Router) Register(router *gin.Engine) {
	router.Static("/static", "./public/static")

	v1 := router.Group("/api")
	{
		r.SysRouteGroup.RegisterBaseRouteGroup(v1)
		r.FrontendRouteGroup.RegisterBaseRouteGroup(v1)
	}
}

var GRouter = new(Router)
