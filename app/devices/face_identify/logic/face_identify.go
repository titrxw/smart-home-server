package logic

import (
	"context"
	"errors"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/model"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/repository"
	"github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	pkglogic "github.com/titrxw/smart-home-server/app/pkg/logic"
	pkgmodel "github.com/titrxw/smart-home-server/app/pkg/model"
	pkgrepository "github.com/titrxw/smart-home-server/app/pkg/repository"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
	"time"
)

const DeviceOperateAddModel = "add_face_model"
const DeviceOperateDelModel = "del_face_model"
const DeviceIdentifyReport = "identify"
const OperateAddModelSettingMinImgLength = 8

type FaceIdentify struct {
	pkglogic.Abstract
}

func (l FaceIdentify) AddDeviceFaceModel(ctx context.Context, userId uint, deviceAppId string, userName string, urls model.FaceUrls) (*model.FaceModel, error) {
	faceModel := &model.FaceModel{
		UserId:      userId,
		DeviceAppId: deviceAppId,
		UserName:    userName,
		Urls:        urls,
		Status:      model.FaceModelStatusDisable,
		CreatedAt:   pkgmodel.LocalTime(time.Now()),
	}

	if !repository.FaceIdentifyDeviceRepository.FaceModel.SaveFaceModel(l.GetDefaultDb(), faceModel) {
		return nil, exception.NewResponseError("添加模型失败")
	}

	err := device.TriggerOperate(ctx,
		facade.GetConfig().GetString("gateway.operate_trigger_url"),
		facade.GetConfig().GetString("face_identify.appid"),
		facade.GetConfig().GetString("face_identify.appsecret"),
		userId,
		deviceAppId,
		DeviceOperateAddModel,
		map[string]interface{}{
			"user_name": userName,
			"urls":      urls,
			"label":     faceModel.ID,
		},
	)
	if err != nil {
		return nil, err
	}

	return faceModel, nil
}

func (l FaceIdentify) DelDeviceFaceModel(ctx context.Context, userId uint, deviceAppId string, label uint) error {
	faceModel := l.GetByUserAndLabel(userId, label)
	if faceModel == nil {
		return errors.New("模型不存在")
	}
	if faceModel.DeviceAppId != deviceAppId {
		return errors.New("模型不存在")
	}

	err := device.TriggerOperate(ctx,
		facade.GetConfig().GetString("gateway.operate_trigger_url"),
		facade.GetConfig().GetString("face_identify.appid"),
		facade.GetConfig().GetString("face_identify.appsecret"),
		userId,
		deviceAppId,
		DeviceOperateDelModel,
		map[string]interface{}{
			"label": faceModel.ID,
		},
	)
	if err == nil {
		faceModel.Status = model.FaceModelStatusDisable
		repository.FaceIdentifyDeviceRepository.FaceModel.SaveFaceModel(l.GetDefaultDb(), faceModel)
	}

	return err
}

func (l FaceIdentify) GetByUserAndLabel(userId uint, label uint) *model.FaceModel {
	return repository.FaceIdentifyDeviceRepository.FaceModel.GetByUserAndLabel(l.GetDefaultDb(), userId, label)
}

func (l FaceIdentify) GetByDeviceAppIdAndLabel(deviceAppId string, label uint) *model.FaceModel {
	return repository.FaceIdentifyDeviceRepository.FaceModel.GetByDeviceAppIdAndLabel(l.GetDefaultDb(), deviceAppId, label)
}

func (l FaceIdentify) GetDeviceFaceModels(userId uint, deviceAppId string, page uint, pageSize uint) *pkgrepository.PageModel {
	return repository.FaceIdentifyDeviceRepository.GetDeviceFaceModels(l.GetDefaultDb(), userId, deviceAppId, page, pageSize)
}

func (l FaceIdentify) GetDeviceFaceModel(userId uint, deviceAppId string, label uint) (*model.FaceModel, error) {
	faceModel := l.GetByUserAndLabel(userId, label)
	if faceModel == nil {
		return nil, exception.NewResponseError("模型不存在")
	}

	if faceModel.DeviceAppId != deviceAppId {
		return nil, exception.NewResponseError("非法操作")
	}

	return faceModel, nil
}

func (l FaceIdentify) OperateReport(gatewayDeviceAppId string, deviceAppId string, message *device.OperateMessage) error {
	if message.EventType == DeviceOperateAddModel {
		result, err := device.IsSuccessResponse(message.Payload)
		if err != nil {
			return err
		}

		if result {
			faceModel := l.GetByDeviceAppIdAndLabel(deviceAppId, uint(message.Payload["label"].(int64)))
			if faceModel == nil {
				return exception.NewResponseError("模型不存在")
			}

			faceModel.Status = model.FaceModelStatusEnable
			repository.FaceIdentifyDeviceRepository.FaceModel.SaveFaceModel(l.GetDefaultDb(), faceModel)
		}
	}

	if message.EventType == DeviceIdentifyReport {
		if _, ok := message.Payload["label"]; !ok {
			return exception.NewResponseError("label 参数缺失")
		}
		if _, ok := message.Payload["label"].(int64); !ok {
			return exception.NewResponseError("label 参数格式错误")
		}
		if _, ok := message.Payload["mat"]; !ok {
			return exception.NewResponseError("mat 参数缺失")
		}
		if _, ok := message.Payload["mat"].(string); !ok {
			return exception.NewResponseError("mat 参数格式错误")
		}

		faceModel := l.GetByDeviceAppIdAndLabel(deviceAppId, uint(message.Payload["label"].(int64)))
		if faceModel == nil {
			return exception.NewResponseError("模型不存在")
		}
		if !faceModel.IsEnable() {
			return exception.NewResponseError("模型不可用")
		}

		//发邮件

		//return a.sendFaceModelIdentifyEmail(user, device, faceModel, message.Payload["mat"].(string))
	}
	return nil
}
