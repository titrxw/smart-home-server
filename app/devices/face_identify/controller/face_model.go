package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/logic"
	"github.com/titrxw/smart-home-server/app/pkg/device"
)

type AddFaceModelRequest struct {
	DeviceAppId string   `form:"device_appid" binding:"required"`
	UserName    string   `form:"user_name" binding:"required"`
	faceUrls    []string `form:"urls" binding:"required,gt=8"`
}

type DelFaceModelRequest struct {
	DeviceAppId string `form:"device_appid" binding:"required"`
	Label       uint   `form:"label" binding:"required"`
}

type DeviceFaceModelDetailRequest struct {
	DeviceAppId string `uri:"device_appid" binding:"required"`
	Label       uint   `uri:"label" binding:"required"`
}

type DeviceFaceModelPageRequest struct {
	Page        uint   `form:"page" binding:"required,page"`
	PageSize    uint   `form:"page_size" binding:"required,page"`
	DeviceAppId string `form:"device_appid" binding:"required"`
}

type DeviceOperateReportRequest struct {
	GatewayDeviceAppId string                 `json:"gateway_appid" form:"gateway_appid" binding:"required"`
	DeviceAppId        string                 `json:"device_appid" form:"device_appid" binding:"required"`
	Message            *device.OperateMessage `json:"message" form:"message" binding:"required"`
}

type FaceModel struct {
	frontend.Abstract
}

func (c FaceModel) AddFaceModel(ctx *gin.Context) {
	addFaceModelRequest := AddFaceModelRequest{}
	if !c.Validate(ctx, &addFaceModelRequest) {
		return
	}

	faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentify.AddDeviceFaceModel(ctx, uint(c.GetUserId(ctx)), addFaceModelRequest.DeviceAppId, addFaceModelRequest.UserName, addFaceModelRequest.faceUrls)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, faceModel)
}

func (c FaceModel) DelFaceModel(ctx *gin.Context) {
	delFaceModelRequest := DelFaceModelRequest{}
	if !c.Validate(ctx, &delFaceModelRequest) {
		return
	}

	err := logic.FaceIdentifyDeviceLogic.FaceIdentify.DelDeviceFaceModel(ctx, uint(c.GetUserId(ctx)), delFaceModelRequest.DeviceAppId, delFaceModelRequest.Label)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}

func (c FaceModel) GetDeviceFaceModelDetail(ctx *gin.Context) {
	deviceFaceModelDetailRequest := DeviceFaceModelDetailRequest{}
	if !c.Validate(ctx, &deviceFaceModelDetailRequest) {
		return
	}

	faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentify.GetDeviceFaceModel(uint(c.GetUserId(ctx)), deviceFaceModelDetailRequest.DeviceAppId, deviceFaceModelDetailRequest.Label)
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

	pageData := logic.FaceIdentifyDeviceLogic.FaceIdentify.GetDeviceFaceModels(uint(c.GetUserId(ctx)), deviceFaceModelPageRequest.DeviceAppId, deviceFaceModelPageRequest.Page, deviceFaceModelPageRequest.PageSize)

	c.JsonResponseWithoutError(ctx, pageData)
}

func (c FaceModel) OperateReport(ctx *gin.Context) {
	deviceOperateReportRequest := DeviceOperateReportRequest{}
	body := ctx.PostForm("body")
	err := json.Unmarshal([]byte(body), &deviceOperateReportRequest)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}
	err = binding.Validator.ValidateStruct(&deviceOperateReportRequest)
	if err != nil {
		c.JsonResponseWithServerError(ctx, c.TranslateValidationError(err))
		return
	}

	err = logic.FaceIdentifyDeviceLogic.FaceIdentify.OperateReport(deviceOperateReportRequest.GatewayDeviceAppId, deviceOperateReportRequest.DeviceAppId, deviceOperateReportRequest.Message)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}
