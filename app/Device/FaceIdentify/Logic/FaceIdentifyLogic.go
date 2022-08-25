package logic

import (
	model2 "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Model"
	repository "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Repository"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository2 "github.com/titrxw/smart-home-server/app/Repository"
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
		return nil, exception.NewLogicError("添加模型失败")
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

func (faceIdentifyLogic FaceIdentifyLogic) GetDeviceFaceModels(device *model.Device, page uint, pageSize uint) *repository2.PageModel {
	return repository.FaceIdentifyDeviceRepository.GetDeviceFaceModels(faceIdentifyLogic.GetDefaultDb(), device.ID, page, pageSize)
}

func (faceIdentifyLogic FaceIdentifyLogic) GetDeviceFaceModel(device *model.Device, label uint) (*model2.FaceModel, error) {
	faceModel := faceIdentifyLogic.GetByLabel(label)
	if faceModel == nil {
		return nil, exception.NewLogicError("模型不存在")
	}

	if faceModel.DeviceId != device.ID {
		return nil, exception.NewLogicError("非法操作")
	}

	return faceModel, nil
}
