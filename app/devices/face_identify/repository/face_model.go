package repository

import (
	"github.com/titrxw/smart-home-server/app/devices/face_identify/model"
	"github.com/titrxw/smart-home-server/app/pkg/repository"
	"gorm.io/gorm"
)

type FaceModel struct {
	repository.Abstract
}

func (r FaceModel) AddFaceModel(db *gorm.DB, faceModel *model.FaceModel) bool {
	result := db.Save(faceModel)

	return result.Error == nil
}

func (r FaceModel) SaveFaceModel(db *gorm.DB, faceModel *model.FaceModel) bool {
	result := db.Save(faceModel)

	return result.Error == nil
}

func (r FaceModel) UpdateFaceModel(db *gorm.DB, faceModel *model.FaceModel) bool {
	result := db.Save(faceModel)

	return result.Error == nil
}

func (r FaceModel) GetByUserAndLabel(db *gorm.DB, userId uint, label uint) *model.FaceModel {
	faceModel := new(model.FaceModel)
	result := db.Where("id = ?", label).Where("user_id = ?", userId).First(faceModel)
	if result.RowsAffected == 1 {
		return faceModel
	}

	return nil
}

func (r FaceModel) GetByDeviceAppIdAndLabel(db *gorm.DB, deviceAppId string, label uint) *model.FaceModel {
	faceModel := new(model.FaceModel)
	result := db.Where("id = ?", label).Where("device_appid = ?", deviceAppId).First(faceModel)
	if result.RowsAffected == 1 {
		return faceModel
	}

	return nil
}

func (r FaceModel) GetDeviceFaceModels(db *gorm.DB, userId uint, deviceAppId string, page uint, pageSize uint) *repository.PageModel {
	faceModels := make([]model.FaceModel, 0)
	pageData := &repository.PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &faceModels,
	}

	totalQuery := db.Model(&model.FaceModel{}).Where("user_id = ?", userId).Where("device_appid = ?", deviceAppId).Where("status != ?", model.FaceModelStatusDisable)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&faceModels)
		pageData.Total = uint64(total)
	}

	return pageData
}
