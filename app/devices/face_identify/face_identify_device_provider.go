package face_identify

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/controller"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/logic"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/middleware"
	"github.com/titrxw/smart-home-server/app/pkg/device"
	pkgmiddleware "github.com/titrxw/smart-home-server/app/pkg/middleware"
	httpserver "github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type Provider struct {
}

func (p Provider) Register(server *httpserver.Server, config *viper.Viper) {
	deviceConfig := device.Device{
		Type:           device.DeviceAppType,
		TypeName:       "face_identify",
		Name:           "识别",
		NeedGateway:    false,
		SupportOperate: []string{logic.DeviceOperateAddModel, logic.DeviceOperateDelModel},
		OperateDesc:    map[string]string{logic.DeviceOperateAddModel: "添加模型", logic.DeviceOperateDelModel: "删除模型"},
		SupportReport:  []string{logic.DeviceIdentifyReport},
		Setting: map[string]interface{}{
			logic.DeviceOperateAddModel: map[string]interface{}{
				"min_img_length": logic.OperateAddModelSettingMinImgLength,
			},
			"report_http_url":  config.GetString("face_identify.server") + "/api/v1/device/face_identify/operate-report",
			"report_appid":     config.GetString("face_identify.appid"),
			"report_appsecret": config.GetString("face_identify.appsecret"),
		},
	}
	manager.RegisterDevice(deviceConfig)

	server.RegisterRouters(func(engine *gin.Engine) {
		console := engine.Group("/console/v1/frontend/device/face_identify", middleware.Oauth{}.Process)
		{
			console.POST("/add", controller.FaceModel{}.AddFaceModel)
			console.POST("/del", controller.FaceModel{}.DelFaceModel)
			console.POST("/list", controller.FaceModel{}.GetDeviceFaceModels)
			console.GET("/detail/:device_appid/:label", controller.FaceModel{}.GetDeviceFaceModelDetail)
		}
		api := engine.Group("/api/v1/device/face_identify", pkgmiddleware.SignCheck{
			AppId:     config.GetString("face_identify.appid"),
			AppSecret: config.GetString("face_identify.appsecret"),
		}.Process)
		{
			api.POST("/operate-report", controller.FaceModel{}.OperateReport)
		}
	})
}
