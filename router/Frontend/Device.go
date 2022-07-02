package frontend

import (
	"github.com/gin-gonic/gin"
	device "github.com/titrxw/smart-home-server/app/Controller/Frontend/Device"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
)

type Device struct {
}

func (devicer *Device) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/device")
	{
		v2.GET("/setting", new(device.DeviceController).DeviceSetting)
		v3 := v2.Group("", new(middleware.OauthMiddleware).Process)
		{
			v3.POST("/add", new(device.DeviceController).AddUserDevice)
			v3.GET("/detail/:device_id", new(device.DeviceController).UserDeviceDetail)
			v3.POST("/update", new(device.DeviceController).UpdateUserDevice)
			v3.POST("/list", new(device.DeviceController).UserDevices)
		}
	}
}
