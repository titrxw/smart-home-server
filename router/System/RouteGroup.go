package system

import "github.com/gin-gonic/gin"

type RouteGroup struct {
	System
}

func (this *RouteGroup) RegisterBaseRouteGroup(router *gin.Engine) {
	this.System.registerRoute(router)
}
