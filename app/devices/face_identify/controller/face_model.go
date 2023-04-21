package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	deviceManagerLogic "github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/logic"
)

type DeviceFaceModelDetailRequest struct {
	DeviceId    uint `uri:"device_id" binding:"required,id"`
	FaceModelId uint `uri:"face_model_id" binding:"required,id"`
}

type DeviceFaceModelPageRequest struct {
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
	DeviceId uint `form:"device_id" binding:"required,id"`
}

type FaceModel struct {
	frontend.Abstract
}

func (c FaceModel) GetDeviceFaceModelDetail(ctx *gin.Context) {
	deviceFaceModelDetailRequest := DeviceFaceModelDetailRequest{}
	if !c.Validate(ctx, &deviceFaceModelDetailRequest) {
		return
	}

	device, err := deviceManagerLogic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceFaceModelDetailRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentify.GetDeviceFaceModel(device, deviceFaceModelDetailRequest.FaceModelId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, faceModel)
}

func (c FaceModel) GetDeviceFaceModels(ctx *gin.Context) {
	deviceFaceModelPageRequest := DeviceFaceModelPageRequest{}
	if !c.Validate(ctx, &deviceFaceModelPageRequest) {
		return
	}

	device, err := deviceManagerLogic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceFaceModelPageRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic.FaceIdentifyDeviceLogic.FaceIdentify.GetDeviceFaceModels(device, deviceFaceModelPageRequest.Page, deviceFaceModelPageRequest.PageSize)

	c.JsonResponseWithoutError(ctx, pageData)
}
