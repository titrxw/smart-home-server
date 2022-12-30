package device

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceGatewayController struct {
	frontend.ControllerAbstract
}

type DeviceGatewayAddRequest struct {
	base.RequestAbstract
	DeviceId        uint `form:"device_id" binding:"required,id"`
	DeviceGatewayId uint `form:"device_gateway_id" binding:"required,id"`
}

func (deviceGatewayController DeviceGatewayController) AddUserGatewayDevice(ctx *gin.Context) {
	deviceGatewayAddRequest := DeviceGatewayAddRequest{}
	if !deviceGatewayController.ValidateFormPost(ctx, &deviceGatewayAddRequest) {
		return
	}

	err := logic.Logic.DeviceGatewayLogic.UserGatewayAddDevice(ctx.Request.Context(), deviceGatewayController.GetUserId(ctx), deviceGatewayAddRequest.DeviceGatewayId, deviceGatewayAddRequest.DeviceId)
	if err != nil {
		deviceGatewayController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceGatewayController.JsonSuccessResponse(ctx)
}
