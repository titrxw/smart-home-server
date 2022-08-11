package logic

import (
	"errors"
	model2 "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Model"
	repository "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Repository"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"time"
)

type FaceIdentifyLogic struct {
	logic.LogicAbstract
}

func (faceIdentifyLogic FaceIdentifyLogic) AddDeviceFaceModel(device *model.Device, userName string, urls model2.FaceUrls) (*model2.FaceModel, error) {
	faceModel := &model2.FaceModel{
		DeviceId:  device.ID,
		UserName:  userName,
		Urls:      urls,
		Status:    model2.FACE_MODEL_STATUS_DISABLE,
		CreatedAt: model.LocalTime(time.Now()),
	}

	if !repository.FaceIdentifyDeviceRepository.FaceModelRepository.AddFaceModel(faceIdentifyLogic.GetDefaultDb(), faceModel) {
		return nil, errors.New("添加模型失败")
	}

	return faceModel, nil
}

func (faceIdentifyLogic FaceIdentifyLogic) GetByLabel(label uint) *model2.FaceModel {
	return repository.FaceIdentifyDeviceRepository.FaceModelRepository.GetByLabel(faceIdentifyLogic.GetDefaultDb(), label)
}

func (faceIdentifyLogic FaceIdentifyLogic) UpdateFaceModelStatus(device *model.Device, label uint, status uint8) bool {
	faceModel := repository.FaceIdentifyDeviceRepository.FaceModelRepository.GetByLabel(faceIdentifyLogic.GetDefaultDb(), label)
	if faceModel == nil {
		return false
	}
	if faceModel.DeviceId != device.ID {
		return false
	}

	faceModel.Status = status
	return repository.FaceIdentifyDeviceRepository.FaceModelRepository.UpdateFaceModel(faceIdentifyLogic.GetDefaultDb(), faceModel)
}
