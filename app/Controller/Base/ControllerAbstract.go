package base

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	controller "github.com/titrxw/go-framework/src/Core/Controller"
	global "github.com/titrxw/go-framework/src/Global"
)

type RequestAbstract struct {
	TimeStamp uint   `json:"time_stamp" form:"time_stamp"`
	Nonce     string `json:"nonce"`
}

type ControllerAbstract struct {
	controller.ControllerAbstract
}

func (controller ControllerAbstract) translateValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		errStr := ""
		for _, e := range validationErrors {
			errStr += e.Translate(global.FApp.Translator) + ";"
		}

		return errStr
	}
}

func (controller ControllerAbstract) ValidateFormPost(ctx *gin.Context, request interface{}) bool {
	err := ctx.ShouldBind(request)
	if err != nil {
		controller.JsonResponseWithServerError(ctx, controller.translateValidationError(err))
		return false
	}

	return true
}

func (controller ControllerAbstract) ValidateQuery(ctx *gin.Context, request interface{}) bool {
	err := ctx.BindQuery(request)
	if err != nil {
		controller.JsonResponseWithServerError(ctx, controller.translateValidationError(err))
		return false
	}

	return true
}

func (controller ControllerAbstract) ValidateFromUri(ctx *gin.Context, request interface{}) bool {
	err := ctx.BindUri(request)
	if err != nil {
		controller.JsonResponseWithServerError(ctx, controller.translateValidationError(err))
		return false
	}

	return true
}
