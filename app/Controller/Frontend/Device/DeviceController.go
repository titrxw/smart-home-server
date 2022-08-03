package device

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceController struct {
	frontend.ControllerAbstract
}

type DeviceAddRequest struct {
	base.RequestAbstract
	DeviceName string `form:"device_name" binding:"required,device_name"`
	DeviceType string `form:"device_type" binding:"required,device_type"`
}

type DeviceUpdateRequest struct {
	base.RequestAbstract
	DeviceId     uint   `form:"device_id" binding:"required,id"`
	DeviceName   string `form:"device_name" binding:"device_name"`
	DeviceStatus uint8  `form:"device_status" binding:"1|2|3"`
}

type DeviceDetailRequest struct {
	base.RequestAbstract
	DeviceId uint `uri:"device_id" binding:"required,id"`
}

type DevicePageRequest struct {
	base.RequestAbstract
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

func (deviceController DeviceController) DeviceSetting(ctx *gin.Context) {
	deviceController.JsonResponseWithoutError(ctx, logic.Logic.DeviceLogic.GetDeviceSupportMap())
}

func (deviceController DeviceController) AddUserDevice(ctx *gin.Context) {
	deviceAddRequest := DeviceAddRequest{}
	if !deviceController.ValidateFormPost(ctx, &deviceAddRequest) {
		return
	}

	if words := logic.Logic.SysSensitiveWordsLogic.GetSensitiveWord(deviceAddRequest.DeviceName); len(words) > 0 {
		deviceController.JsonResponseWithServerError(ctx, errors.New("设备名包含敏感字符 "+strings.Join(words, ",")))
		return
	}

	device, err := logic.Logic.DeviceLogic.CreateUserDevice(ctx, deviceController.GetUserId(ctx), deviceAddRequest.DeviceName, deviceAddRequest.DeviceType)
	if err != nil {
		deviceController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceController.JsonResponseWithoutError(ctx, device)
}

func (deviceController DeviceController) UserDevices(ctx *gin.Context) {
	devicePageRequest := DevicePageRequest{}
	if !deviceController.ValidateFormPost(ctx, &devicePageRequest) {
		return
	}

	pageData := logic.Logic.DeviceLogic.GetUserDevices(deviceController.GetUserId(ctx), devicePageRequest.Page, devicePageRequest.PageSize)

	deviceController.JsonResponseWithoutError(ctx, pageData)
}

func (deviceController DeviceController) UserDeviceDetail(ctx *gin.Context) {
	deviceDetailRequest := DeviceDetailRequest{}
	if !deviceController.ValidateFromUri(ctx, &deviceDetailRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceController.GetUserId(ctx), deviceDetailRequest.DeviceId)
	if err != nil {
		deviceController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceController.JsonResponseWithoutError(ctx, device)
}

func (deviceController DeviceController) UpdateUserDevice(ctx *gin.Context) {
	deviceUpdateRequest := DeviceUpdateRequest{}
	if !deviceController.ValidateFormPost(ctx, &deviceUpdateRequest) {
		return
	}

	if deviceUpdateRequest.DeviceName == "" || deviceUpdateRequest.DeviceStatus == 0 {
		deviceController.JsonResponseWithServerError(ctx, "参数错误")
		return
	}
	if deviceUpdateRequest.DeviceName != "" {
		if words := logic.Logic.SysSensitiveWordsLogic.GetSensitiveWord(deviceUpdateRequest.DeviceName); len(words) > 0 {
			deviceController.JsonResponseWithServerError(ctx, errors.New("设备名包含敏感字符 "+strings.Join(words, ",")))
			return
		}
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceController.GetUserId(ctx), deviceUpdateRequest.DeviceId)
	if err != nil {
		deviceController.JsonResponseWithServerError(ctx, err)
		return
	}

	if deviceUpdateRequest.DeviceName != "" {
		device.Name = deviceUpdateRequest.DeviceName
	}
	if deviceUpdateRequest.DeviceStatus != 0 {
		device.DeviceStatus = deviceUpdateRequest.DeviceStatus
	}
	err = logic.Logic.DeviceLogic.UpdateDevice(ctx, device)
	if err != nil {
		deviceController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceController.JsonSuccessResponse(ctx)
}
