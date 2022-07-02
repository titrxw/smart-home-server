package frontend

import (
	"github.com/gin-gonic/gin"
	device "github.com/titrxw/smart-home-server/app/Controller/Frontend/Device"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
)

type DeviceOperate struct {
}

func (deviceOperate *DeviceOperate) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/device-operate", new(middleware.OauthMiddleware).Process)
	{
		v2.POST("/trigger", new(device.DeviceOperateLogController).TriggerOperate)
		v2.POST("/detail", new(device.DeviceOperateLogController).OperateDetail)
		v2.POST("/list", new(device.DeviceOperateLogController).DeviceOperateLog)
	}
}
