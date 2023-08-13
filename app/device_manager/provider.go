package device_manager

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	deviceconsole "github.com/titrxw/smart-home-server/app/device_manager/console"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/api"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/backend"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/device"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/user"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/system/util"
	middleware2 "github.com/titrxw/smart-home-server/app/device_manager/middleware"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/middleware"
	"github.com/we7coreteam/w7-rangine-go-support/src/console"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
	httpserver "github.com/we7coreteam/w7-rangine-go/src/http/server"
	"regexp"
	"unicode/utf8"
)

type Provider struct {
}

func (p Provider) Register(httpServer *httpserver.Server, consoleManager console.Console) {
	p.RegisterValidateRule()
	p.RegisterHttpRoutes(httpServer)

	consoleManager.RegisterCommand(&deviceconsole.CreateAppCommand{})
}

func (p Provider) RegisterValidateRule() {
	if v, ok := facade.GetValidator().Engine().(*validator.Validate); ok {
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

func (p Provider) RegisterHttpRoutes(server *httpserver.Server) {
	server.RegisterRouters(func(engine *gin.Engine) {
		console := engine.Group("/console/v1")
		{
			utilGroup := console.Group("/util")
			{
				utilGroup.GET("/captcha", util.Captcha{}.Captcha)
				utilGroup.POST("/attach/upload/image", middleware2.AppCheck{}.Process, util.Upload{}.UploadImage)
			}

			userFrontendGroup := console.Group("/frontend/user")
			{
				oauthGroup := userFrontendGroup.Group("/oauth")
				{
					oauthGroup.POST("/register-email", middleware.Captcha{}.Process, user.Oauth{}.SendRegisterEmailCode)
					oauthGroup.POST("/register", user.Oauth{}.Register)
					oauthGroup.POST("/login", middleware.Captcha{}.Process, user.Oauth{}.Login)
					oauthGroup.GET("/logout", middleware.Oauth{}.Process, user.Oauth{}.Logout)
					oauthGroup.GET("/info", middleware.Oauth{}.Process, user.Info{}.Info)
				}

				deviceGroup := userFrontendGroup.Group("/device")
				{
					deviceGroup.GET("/setting", device.Device{}.DeviceSetting)

					deviceManagerGroup := deviceGroup.Group("", middleware.Oauth{}.Process)
					{
						deviceManagerGroup.POST("/add", device.Device{}.AddUserDevice)
						deviceManagerGroup.GET("/detail/:device_id", device.Device{}.UserDeviceDetail)
						deviceManagerGroup.POST("/update", device.Device{}.UpdateUserDevice)
						deviceManagerGroup.POST("/list", device.Device{}.UserDevices)
						deviceManagerGroup.POST("/gateway/bind", device.Gateway{}.UserDeviceBindGateway)

						deviceOperateGroup := deviceManagerGroup.Group("/operate")
						{
							deviceOperateGroup.POST("/list", device.Operate{}.DeviceOperateLog)
							deviceOperateGroup.POST("/detail", device.Operate{}.OperateDetail)
						}

						deviceReportGroup := deviceManagerGroup.Group("/report")
						{
							deviceReportGroup.POST("/list", device.ReportLog{}.DeviceReportLog)
							deviceReportGroup.POST("/detail", device.ReportLog{}.ReportDetail)
						}
					}
				}
			}

			backendGroup := console.Group("/backend")
			{
				deviceApiGroup := backendGroup.Group("/device")
				{
					deviceApiGroup.POST("/register", middleware.Oauth{}.Process, backend.Device{}.AddSupportDevice)
				}
			}
		}

		apiGroup := engine.Group("/api/v1", middleware2.AppCheck{}.Process)
		{
			deviceApiGroup := apiGroup.Group("/device")
			{
				deviceApiGroup.POST("/trigger-operate", api.Device{}.TriggerOperate)
			}
		}
	})
}
