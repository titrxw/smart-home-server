package system

import "github.com/gin-gonic/gin"

type SysRouteGroup struct {
	System
	Util
}

func (this *SysRouteGroup) RegisterBaseRouteGroup(router *gin.RouterGroup) {
	this.System.registerRoute(router)
	this.Util.registerRoute(router)
}
