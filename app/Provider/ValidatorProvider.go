package provider

import (
	ut "github.com/go-playground/universal-translator"
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	global "github.com/titrxw/go-framework/src/Global"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"regexp"
	"unicode/utf8"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidatorProvider struct {
	provider.ProviderAbstract
}

func (this *ValidatorProvider) Register(options interface{}) {
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
		v.RegisterTranslation("page", global.FApp.Translator, func(ut ut.Translator) error {
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
		v.RegisterTranslation("device_name", global.FApp.Translator, func(ut ut.Translator) error {
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

				if _, ok = logic.Logic.DeviceLogic.GetDeviceSupportMap()[deviceType]; ok {
					return true
				}

				return false
			}

			return false
		})
		v.RegisterTranslation("device_type", global.FApp.Translator, func(ut ut.Translator) error {
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
		v.RegisterTranslation("user_name", global.FApp.Translator, func(ut ut.Translator) error {
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
		v.RegisterTranslation("mobile", global.FApp.Translator, func(ut ut.Translator) error {
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
		v.RegisterTranslation("password", global.FApp.Translator, func(ut ut.Translator) error {
			return ut.Add("password", "{0} 格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		})
	}

}
