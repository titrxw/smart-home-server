package device

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
)

type OperateLog struct {
	frontend.Abstract
}

type OperateRequest struct {
	DeviceId       uint                    `form:"device_id" binding:"required,id"`
	OperateType    model.DeviceOperateType `form:"operate_type" binding:"required"`
	OperatePayload model.OperatePayload    `form:"operate_payload" binding:"required"`
}

type OperateDetailRequest struct {
	DeviceId      uint   `form:"device_id" binding:"required,id"`
	OperateNumber string `form:"operate_number" binding:"required"`
}

type OperatePageRequest struct {
	DeviceId uint `form:"device_id" binding:"required,id"`
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

func (c OperateLog) TriggerOperate(ctx *gin.Context) {
	deviceOperateRequest := OperateRequest{}
	if !c.Validate(ctx, &deviceOperateRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceOperateRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	operateLog, err := logic.Logic.DeviceOperate.TriggerOperate(ctx.Request.Context(), device, deviceOperateRequest.OperateType, deviceOperateRequest.OperatePayload, 2)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, operateLog)
}

func (c OperateLog) OperateDetail(ctx *gin.Context) {
	deviceOperateDetailRequest := OperateDetailRequest{}
	if !c.Validate(ctx, &deviceOperateDetailRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceOperateDetailRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	operateLog, err := logic.Logic.DeviceOperate.GetDeviceOperateLogByNumber(device, deviceOperateDetailRequest.OperateNumber)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, operateLog)
}

func (c OperateLog) DeviceOperateLog(ctx *gin.Context) {
	deviceOperatePageRequest := OperatePageRequest{}
	if !c.Validate(ctx, &deviceOperatePageRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceOperatePageRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic.Logic.DeviceOperate.GetDeviceOperates(device, deviceOperatePageRequest.Page, deviceOperatePageRequest.PageSize)

	c.JsonResponseWithoutError(ctx, pageData)
}
