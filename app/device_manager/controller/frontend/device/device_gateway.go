package device

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/internal/logic"
)

type Gateway struct {
	frontend.Abstract
}

type GatewayAddRequest struct {
	DeviceId        uint `form:"device_id" binding:"required,id"`
	DeviceGatewayId uint `form:"device_gateway_id" binding:"required,id"`
}

func (c Gateway) UserDeviceBindGateway(ctx *gin.Context) {
	deviceGatewayAddRequest := GatewayAddRequest{}
	if !c.Validate(ctx, &deviceGatewayAddRequest) {
		return
	}

	err := logic.Logic.DeviceGateway.UserGatewayAddDevice(ctx.Request.Context(), c.GetUserId(ctx), deviceGatewayAddRequest.DeviceGatewayId, deviceGatewayAddRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}
