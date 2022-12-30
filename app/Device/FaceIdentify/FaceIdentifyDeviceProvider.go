package faceIdentify

import (
	"github.com/gin-gonic/gin"
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	controller "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Controller"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
	mqtt "github.com/titrxw/smart-home-server/app/Mqtt"
	subscribe "github.com/titrxw/smart-home-server/app/Mqtt/Subscribe"
	"github.com/titrxw/smart-home-server/router"
)

type FaceIdentifyDeviceProvider struct {
	provider.ProviderAbstract
}

func (faceIdentifyDeviceProvider *FaceIdentifyDeviceProvider) Register(options interface{}) {
	adapter := new(FaceIdentifyDeviceAdapter)
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(adapter)

	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe(adapter.GetReportTopic("+")))

	router.GRouter.RegisterRouterProvider(func(engine *gin.Engine) {
		v1 := engine.Group("/api/v1/frontend/device/"+adapter.GetDeviceConfig().TypeName, new(middleware.OauthMiddleware).Process)
		{
			v1.GET("/detail/:device_id/:face_model_id", new(controller.FaceModelController).GetDeviceFaceModelDetail)
			v1.POST("/list", new(controller.FaceModelController).GetDeviceFaceModels)
		}
	})
}
