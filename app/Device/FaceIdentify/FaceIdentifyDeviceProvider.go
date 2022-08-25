package faceIdentify

import (
	"github.com/gin-gonic/gin"
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	controller "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Controller"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
	"github.com/titrxw/smart-home-server/router"
)

type FaceIdentifyDeviceProvider struct {
	provider.ProviderAbstract
}

func (faceIdentifyDeviceProvider *FaceIdentifyDeviceProvider) Register(options interface{}) {
	adapter := new(FaceIdentifyDeviceAdapter)
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(adapter)

	router.GRouter.RegisterRouterProvider(func(engine *gin.Engine) {
		v1 := engine.Group("/api/frontend/device/"+adapter.GetDeviceConfig().Type, new(middleware.OauthMiddleware).Process)
		{
			v1.GET("/detail/:device_id/:face_model_id", new(controller.FaceModelController).GetDeviceFaceModelDetail)
			v1.POST("/list", new(controller.FaceModelController).GetDeviceFaceModels)
		}
	})
}
