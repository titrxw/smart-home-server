package frontend

import "github.com/gin-gonic/gin"

type FrontendRouteGroup struct {
	User
	Device
	DeviceOperate
}

func (frontendRouteGroup *FrontendRouteGroup) RegisterBaseRouteGroup(router *gin.RouterGroup) {
	v1 := router.Group("/frontend")
	{
		v2 := v1.Group("/user")
		{
			frontendRouteGroup.User.registerRoute(v2)
			frontendRouteGroup.Device.registerRoute(v2)
			frontendRouteGroup.DeviceOperate.registerRoute(v2)
		}
	}
}
