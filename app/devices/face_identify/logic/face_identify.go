package logic

import (
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	deviceManagerModel "github.com/titrxw/smart-home-server/app/device_manager/model"
	deviceManagerRepository "github.com/titrxw/smart-home-server/app/device_manager/repository"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/model"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/repository"
	"time"
)

type FaceIdentify struct {
	logic.Abstract
}

func (l FaceIdentify) AddDeviceFaceModel(device *deviceManagerModel.Device, userName string, urls model.FaceUrls) (*model.FaceModel, error) {
	faceModel := &model.FaceModel{
		DeviceId:  device.ID,
		UserName:  userName,
		Urls:      urls,
		Status:    model.FaceModelStatusDisable,
		CreatedAt: deviceManagerModel.LocalTime(time.Now()),
	}

	if !repository.FaceIdentifyDeviceRepository.FaceModel.AddFaceModel(l.GetDefaultDb(), faceModel) {
		return nil, exception.NewResponseError("添加模型失败")
	}

	return faceModel, nil
}

func (l FaceIdentify) GetByLabel(label uint) *model.FaceModel {
	return repository.FaceIdentifyDeviceRepository.FaceModel.GetByLabel(l.GetDefaultDb(), label)
}

func (l FaceIdentify) UpdateFaceModelStatus(device *deviceManagerModel.Device, label uint, status uint8) bool {
	faceModel := repository.FaceIdentifyDeviceRepository.FaceModel.GetByLabel(l.GetDefaultDb(), label)
	if faceModel == nil {
		return false
	}
	if faceModel.DeviceId != device.ID {
		return false
	}

	faceModel.Status = status
	return repository.FaceIdentifyDeviceRepository.FaceModel.UpdateFaceModel(l.GetDefaultDb(), faceModel)
}

func (l FaceIdentify) GetDeviceFaceModels(device *deviceManagerModel.Device, page uint, pageSize uint) *deviceManagerRepository.PageModel {
	return repository.FaceIdentifyDeviceRepository.GetDeviceFaceModels(l.GetDefaultDb(), device.ID, page, pageSize)
}

func (l FaceIdentify) GetDeviceFaceModel(device *deviceManagerModel.Device, label uint) (*model.FaceModel, error) {
	faceModel := l.GetByLabel(label)
	if faceModel == nil {
		return nil, exception.NewResponseError("模型不存在")
	}

	if faceModel.DeviceId != device.ID {
		return nil, exception.NewResponseError("非法操作")
	}

	return faceModel, nil
}
