package device

import (
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

type Device struct {
	frontend.Abstract
}

type AddRequest struct {
	DeviceName string `form:"device_name" binding:"required,device_name"`
	DeviceType string `form:"device_type" binding:"required,device_type"`
}

type UpdateRequest struct {
	DeviceId     uint   `form:"device_id" binding:"required,id"`
	DeviceName   string `form:"device_name" binding:"device_name"`
	DeviceStatus uint8  `form:"device_status" binding:"1|2|3"`
}

type DetailRequest struct {
	DeviceId uint `uri:"device_id" binding:"required,id"`
}

type PageRequest struct {
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

func (c Device) DeviceSetting(ctx *gin.Context) {
	c.JsonResponseWithoutError(ctx, manager.GetDeviceSupportMap())
}

func (c Device) AddUserDevice(ctx *gin.Context) {
	deviceAddRequest := AddRequest{}
	if !c.Validate(ctx, &deviceAddRequest) {
		return
	}

	if words := logic.Logic.SysSensitiveWords.GetSensitiveWord(deviceAddRequest.DeviceName); len(words) > 0 {
		c.JsonResponseWithServerError(ctx, exception.NewResponseError("设备名包含敏感字符 "+strings.Join(words, ",")))
		return
	}

	device, err := logic.Logic.Device.CreateUserDevice(ctx.Request.Context(), c.GetUserId(ctx), deviceAddRequest.DeviceName, deviceAddRequest.DeviceType)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, device)
}

func (c Device) UserDevices(ctx *gin.Context) {
	devicePageRequest := PageRequest{}
	if !c.Validate(ctx, &devicePageRequest) {
		return
	}

	pageData := logic.Logic.Device.GetUserDevices(c.GetUserId(ctx), devicePageRequest.Page, devicePageRequest.PageSize)

	c.JsonResponseWithoutError(ctx, pageData)
}

func (c Device) UserDeviceDetail(ctx *gin.Context) {
	deviceDetailRequest := DetailRequest{}
	if !c.Validate(ctx, &deviceDetailRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceDetailRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}
	if manager.GetDeviceConfigByDeviceType(device.TypeName).NeedGateway {
		device.GatewayDevice = logic.Logic.DeviceGateway.GetGatewayDevice(ctx, device)
	}

	c.JsonResponseWithoutError(ctx, device)
}

func (c Device) UpdateUserDevice(ctx *gin.Context) {
	deviceUpdateRequest := UpdateRequest{}
	if !c.Validate(ctx, &deviceUpdateRequest) {
		return
	}

	if deviceUpdateRequest.DeviceName == "" || deviceUpdateRequest.DeviceStatus == 0 {
		c.JsonResponseWithServerError(ctx, exception.NewResponseError("参数错误"))
		return
	}
	if deviceUpdateRequest.DeviceName != "" {
		if words := logic.Logic.SysSensitiveWords.GetSensitiveWord(deviceUpdateRequest.DeviceName); len(words) > 0 {
			c.JsonResponseWithServerError(ctx, exception.NewResponseError("设备名包含敏感字符 "+strings.Join(words, ",")))
			return
		}
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceUpdateRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	if deviceUpdateRequest.DeviceName != "" {
		device.Name = deviceUpdateRequest.DeviceName
	}
	if deviceUpdateRequest.DeviceStatus != 0 {
		device.DeviceStatus = deviceUpdateRequest.DeviceStatus
	}
	err = logic.Logic.Device.UpdateDevice(ctx.Request.Context(), device)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}
