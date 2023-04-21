package device_manager

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/device"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/user"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/system/util"
	"github.com/titrxw/smart-home-server/app/device_manager/middleware"
	"github.com/titrxw/smart-home-server/app/device_manager/mqtt/subscribe"
	"github.com/titrxw/smart-home-server/app/devices/manager"
	"github.com/titrxw/smart-home-server/app/mqtt"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
	httpserver "github.com/we7coreteam/w7-rangine-go/src/http/server"
	"regexp"
	"unicode/utf8"
)

type Provider struct {
	provider.Abstract
}

func (p Provider) Register() {
	p.RegisterValidateRule()
	p.RegisterHttpRoutes()
	p.RegisterMqttSubscribe()
}

func (p Provider) RegisterValidateRule() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
			if id, ok := fl.Field().Interface().(uint); ok {
				if id > 0 {
					return true
				}
			}

			return false
		})

		v.RegisterValidation("page", func(fl validator.FieldLevel) bool {
			if page, ok := fl.Field().Interface().(uint); ok {
				if page > 0 {
					return true
				}
			}

			return false
		})
		v.RegisterTranslation("page", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("page", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("page", fe.Field())
			return t
		})

		v.RegisterValidation("device_name", func(fl validator.FieldLevel) bool {
			if userName, ok := fl.Field().Interface().(string); ok {
				if utf8.RuneCountInString(userName) > 20 {
					return false
				}
				if !regexp.MustCompile(`^[a-zA-Z0-9\x{4e00}-\x{9fa5}]+$`).MatchString(userName) {
					return false
				}

				return true
			}

			return false
		})
		v.RegisterTranslation("device_name", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("device_name", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("device_name", fe.Field())
			return t
		})

		v.RegisterValidation("device_type", func(fl validator.FieldLevel) bool {
			if deviceType, ok := fl.Field().Interface().(string); ok {
				if utf8.RuneCountInString(deviceType) == 0 {
					return false
				}

				if _, ok = manager.GetDeviceSupportMap()[deviceType]; ok {
					return true
				}

				return false
			}

			return false
		})
		v.RegisterTranslation("device_type", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("device_type", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("device_type", fe.Field())
			return t
		})

		v.RegisterValidation("user_name", func(fl validator.FieldLevel) bool {
			if userName, ok := fl.Field().Interface().(string); ok {
				if utf8.RuneCountInString(userName) > 20 {
					return false
				}
				if !regexp.MustCompile(`^[a-zA-Z0-9\x{4e00}-\x{9fa5}]+$`).MatchString(userName) {
					return false
				}

				return true
			}

			return false
		})
		v.RegisterTranslation("user_name", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("user_name", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("user_name", fe.Field())
			return t
		})

		v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
			if mobile, ok := fl.Field().Interface().(string); ok {
				if regexp.MustCompile(`^1[3456789]\d{9}$`).MatchString(mobile) {
					return true
				}
			}
			return false
		})
		v.RegisterTranslation("mobile", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})

		v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			if password, ok := fl.Field().Interface().(string); ok {
				if len(password) < 8 || len(password) > 18 {
					return false
				}
				if b, err := regexp.MatchString(`[0-9]{1}`, password); !b || err != nil {
					return false
				}
				if b, err := regexp.MatchString(`[a-z]{1}`, password); !b || err != nil {
					return false
				}
				if b, err := regexp.MatchString(`[A-Z]{1}`, password); !b || err != nil {
					return false
				}

				return true
			}

			return false
		})
		v.RegisterTranslation("password", facade.GetTranslator(), func(ut ut.Translator) error {
			return ut.Add("password", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		})
	}
}

func (p Provider) RegisterHttpRoutes() {
	httpserver.RegisterRouters(func(engine *gin.Engine) {
		v1 := engine.Group("/api/v1")
		{
			utilGroup := v1.Group("/util")
			{
				utilGroup.GET("/captcha", util.Captcha{}.Captcha)
				utilGroup.POST("/attach/upload/image", middleware.AppCheck{}.Process, util.Upload{}.UploadImage)
			}

			userGroup := v1.Group("/frontend/user")
			{
				oauth := userGroup.Group("/oauth")
				{
					oauth.POST("/register-email", middleware.Captcha{}.Process, user.Oauth{}.SendRegisterEmailCode)
					oauth.POST("/register", user.Oauth{}.Register)
					oauth.POST("/login", middleware.Captcha{}.Process, user.Oauth{}.Login)
					oauth.GET("/logout", middleware.Oauth{}.Process, user.Oauth{}.Logout)
					oauth.GET("/info", middleware.Oauth{}.Process, user.Info{}.Info)
				}

				deviceGroup := userGroup.Group("/device")
				{
					deviceGroup.GET("/setting", device.Device{}.DeviceSetting)
					deviceManager := deviceGroup.Group("", middleware.Oauth{}.Process)
					{
						deviceManager.POST("/add", device.Device{}.AddUserDevice)
						deviceManager.GET("/detail/:device_id", device.Device{}.UserDeviceDetail)
						deviceManager.POST("/update", device.Device{}.UpdateUserDevice)
						deviceManager.POST("/list", device.Device{}.UserDevices)
					}
				}
				userGroup.POST("/device-gateway/bind", middleware.Oauth{}.Process, device.Gateway{}.AddUserGatewayDevice)

				deviceOperate := userGroup.Group("/device-operate", middleware.Oauth{}.Process)
				{
					deviceOperate.POST("/trigger", device.OperateLog{}.TriggerOperate)
					deviceOperate.POST("/detail", device.OperateLog{}.OperateDetail)
					deviceOperate.POST("/list", device.OperateLog{}.DeviceOperateLog)
				}

				deviceReport := userGroup.Group("/device-report", middleware.Oauth{}.Process)
				{
					deviceReport.POST("/detail", device.ReportLog{}.ReportDetail)
					deviceReport.POST("/list", device.ReportLog{}.DeviceReportLog)
				}
			}
		}
	})
}

func (p Provider) RegisterMqttSubscribe() {
	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe("/iot/" + helper.GetAppName() + "/device/+/report"))
	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceStatusChangeSubscribe("$SYS/brokers/+/clients/+/+"))
}
