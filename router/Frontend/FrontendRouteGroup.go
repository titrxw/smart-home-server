package frontend

import "github.com/gin-gonic/gin"

type FrontendRouteGroup struct {
	User
	Device
	DeviceOperate
}

func (this *FrontendRouteGroup) RegisterBaseRouteGroup(router *gin.RouterGroup) {
	v1 := router.Group("/frontend")
	{
		v2 := v1.Group("/user")
		{
			this.User.registerRoute(v2)
			this.Device.registerRoute(v2)
			this.DeviceOperate.registerRoute(v2)
		}
	}
}
