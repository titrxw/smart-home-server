package frontend

import (
	"github.com/gin-gonic/gin"
	device "github.com/titrxw/smart-home-server/app/Controller/Frontend/Device"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
)

type DeviceReport struct {
}

func (deviceReport *DeviceReport) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/device-report", new(middleware.OauthMiddleware).Process)
	{
		v2.POST("/detail", new(device.DeviceReportLogController).ReportDetail)
		v2.POST("/list", new(device.DeviceReportLogController).DeviceReportLog)
	}
}
