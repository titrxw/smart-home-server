package device

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceOperateLogController struct {
	frontend.ControllerAbstract
}

type DeviceOperateRequest struct {
	base.RequestAbstract
	DeviceId       uint                    `form:"device_id" binding:"required,id"`
	OperateType    model.DeviceOperateType `form:"operate_type" binding:"required"`
	OperatePayload model.OperatePayload    `form:"operate_payload" binding:"required"`
}

type DeviceOperateDetailRequest struct {
	base.RequestAbstract
	DeviceId      uint   `form:"device_id" binding:"required,id"`
	OperateNumber string `form:"operate_number" binding:"required"`
}

type DeviceOperatePageRequest struct {
	base.RequestAbstract
	DeviceId uint `form:"device_id" binding:"required,id"`
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

func (deviceOperateController DeviceOperateLogController) TriggerOperate(ctx *gin.Context) {
	deviceOperateRequest := DeviceOperateRequest{}
	if !deviceOperateController.ValidateFormPost(ctx, &deviceOperateRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceOperateController.GetUserId(ctx), deviceOperateRequest.DeviceId)
	if err != nil {
		deviceOperateController.JsonResponseWithServerError(ctx, err)
		return
	}

	operateLog, err := logic.Logic.DeviceOperateLogic.TriggerOperate(ctx.Request.Context(), device, deviceOperateRequest.OperateType, deviceOperateRequest.OperatePayload, 2)
	if err != nil {
		deviceOperateController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceOperateController.JsonResponseWithoutError(ctx, operateLog)
}

func (deviceOperateController DeviceOperateLogController) OperateDetail(ctx *gin.Context) {
	deviceOperateDetailRequest := DeviceOperateDetailRequest{}
	if !deviceOperateController.ValidateFormPost(ctx, &deviceOperateDetailRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceOperateController.GetUserId(ctx), deviceOperateDetailRequest.DeviceId)
	if err != nil {
		deviceOperateController.JsonResponseWithServerError(ctx, err)
		return
	}

	operateLog, err := logic.Logic.DeviceOperateLogic.GetDeviceOperateLogByNumber(device, deviceOperateDetailRequest.OperateNumber)
	if err != nil {
		deviceOperateController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceOperateController.JsonResponseWithoutError(ctx, operateLog)
}

func (deviceOperateController DeviceOperateLogController) DeviceOperateLog(ctx *gin.Context) {
	deviceOperatePageRequest := DeviceOperatePageRequest{}
	if !deviceOperateController.ValidateFormPost(ctx, &deviceOperatePageRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceOperateController.GetUserId(ctx), deviceOperatePageRequest.DeviceId)
	if err != nil {
		deviceOperateController.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic.Logic.DeviceOperateLogic.GetDeviceOperates(device, deviceOperatePageRequest.Page, deviceOperatePageRequest.PageSize)

	deviceOperateController.JsonResponseWithoutError(ctx, pageData)
}
