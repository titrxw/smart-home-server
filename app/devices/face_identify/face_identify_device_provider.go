package face_identify

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/middleware"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/controller"
	"github.com/titrxw/smart-home-server/app/devices/manager"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
	httpserver "github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type Provider struct {
	provider.Abstract
}

func (p *Provider) Register() {
	adapter := new(DeviceAdapter)
	manager.RegisterDevice(adapter)

	httpserver.RegisterRouters(func(engine *gin.Engine) {
		v1 := engine.Group("/api/v1/frontend/device/"+adapter.GetDeviceConfig().TypeName, middleware.Oauth{}.Process)
		{
			v1.GET("/detail/:device_id/:face_model_id", controller.FaceModel{}.GetDeviceFaceModelDetail)
			v1.POST("/list", controller.FaceModel{}.GetDeviceFaceModels)
		}
	})
}
