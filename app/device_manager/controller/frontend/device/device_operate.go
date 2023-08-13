package device

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/internal/logic"
)

type Operate struct {
	frontend.Abstract
}

type OperatePageRequest struct {
	DeviceId uint `form:"device_id" binding:"required,id"`
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

type OperateDetailRequest struct {
	DeviceId      uint   `form:"device_id" binding:"required,id"`
	OperateNumber string `form:"operate_number" binding:"required"`
}

func (c Operate) DeviceOperateLog(ctx *gin.Context) {
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

func (c Operate) OperateDetail(ctx *gin.Context) {
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
