package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/internal/model"
)

type Device struct {
	frontend.Abstract
}

type OperateRequest struct {
	UserId         model.UID               `json:"user_id" form:"user_id" binding:"required"`
	DeviceAppId    string                  `json:"device_appid" form:"device_appid" binding:"required"`
	OperateType    model.DeviceOperateType `json:"operate_type" form:"operate_type" binding:"required"`
	OperatePayload model.OperatePayload    `json:"operate_payload" form:"operate_payload" binding:"required"`
}

func (c Device) TriggerOperate(ctx *gin.Context) {
	deviceOperateRequest := OperateRequest{}
	body := ctx.PostForm("body")
	err := json.Unmarshal([]byte(body), &deviceOperateRequest)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}
	err = binding.Validator.ValidateStruct(&deviceOperateRequest)
	if err != nil {
		c.JsonResponseWithServerError(ctx, c.TranslateValidationError(err))
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceByDeviceId(deviceOperateRequest.UserId, deviceOperateRequest.DeviceAppId)
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
