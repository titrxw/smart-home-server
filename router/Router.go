package router

import (
	"github.com/gin-gonic/gin"
	frontend "github.com/titrxw/smart-home-server/router/Frontend"
	system "github.com/titrxw/smart-home-server/router/System"
)

type Router struct {
	routerProviders []func(engine *gin.Engine)

	system.SysRouteGroup
	frontend.FrontendRouteGroup
}

func (r *Router) RegisterRouterProvider(routerProvider func(engine *gin.Engine)) {
	r.routerProviders = append(r.routerProviders, routerProvider)
}

func (r *Router) Register(router *gin.Engine) {
	for _, routerProvider := range r.routerProviders {
		routerProvider(router)
	}

	router.Static("/static", "./public/static")
	router.Static("/img", "./public/upload/img")

	v1 := router.Group("/api/v1")
	{
		r.SysRouteGroup.RegisterBaseRouteGroup(v1)
		r.FrontendRouteGroup.RegisterBaseRouteGroup(v1)
	}
}

var GRouter = new(Router)
