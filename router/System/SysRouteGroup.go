package system

import "github.com/gin-gonic/gin"

type SysRouteGroup struct {
	System
	Util
}

func (sysRouteGroup *SysRouteGroup) RegisterBaseRouteGroup(router *gin.RouterGroup) {
	sysRouteGroup.System.registerRoute(router)
	sysRouteGroup.Util.registerRoute(router)
}
