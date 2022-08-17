package controller

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic2 "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Logic"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceFaceModelDetailRequest struct {
	base.RequestAbstract
	DeviceId    uint `uri:"device_id" binding:"required,id"`
	FaceModelId uint `uri:"face_model_id" binding:"required,id"`
}

type DeviceFaceModelPageRequest struct {
	base.RequestAbstract
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
	DeviceId uint `form:"device_id" binding:"required,id"`
}

type FaceModelController struct {
	frontend.ControllerAbstract
}

func (faceModelController FaceModelController) GetDeviceFaceModelDetail(ctx *gin.Context) {
	deviceFaceModelDetailRequest := DeviceFaceModelDetailRequest{}
	if !faceModelController.ValidateFromUri(ctx, &deviceFaceModelDetailRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(faceModelController.GetUserId(ctx), deviceFaceModelDetailRequest.DeviceId)
	if err != nil {
		faceModelController.JsonResponseWithServerError(ctx, err)
		return
	}

	faceModel, err := logic2.FaceIdentifyDeviceLogic.FaceIdentifyLogic.GetDeviceFaceModel(device, deviceFaceModelDetailRequest.FaceModelId)
	if err != nil {
		faceModelController.JsonResponseWithServerError(ctx, err)
		return
	}

	faceModelController.JsonResponseWithoutError(ctx, faceModel)
}

func (faceModelController FaceModelController) GetDeviceFaceModels(ctx *gin.Context) {
	deviceFaceModelPageRequest := DeviceFaceModelPageRequest{}
	if !faceModelController.ValidateFormPost(ctx, &deviceFaceModelPageRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(faceModelController.GetUserId(ctx), deviceFaceModelPageRequest.DeviceId)
	if err != nil {
		faceModelController.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic2.FaceIdentifyDeviceLogic.FaceIdentifyLogic.GetDeviceFaceModels(device, deviceFaceModelPageRequest.Page, deviceFaceModelPageRequest.PageSize)

	faceModelController.JsonResponseWithoutError(ctx, pageData)
}
