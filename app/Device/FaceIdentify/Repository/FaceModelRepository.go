package repository

import (
	model "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"gorm.io/gorm"
)

type FaceModelRepository struct {
	repository.RepositoryAbstract
}

func (faceModelRepository FaceModelRepository) AddFaceModel(db *gorm.DB, faceModel *model.FaceModel) bool {
	result := db.Save(faceModel)

	return result.Error == nil
}

func (faceModelRepository FaceModelRepository) UpdateFaceModel(db *gorm.DB, faceModel *model.FaceModel) bool {
	result := db.Save(faceModel)

	return result.Error == nil
}

func (faceModelRepository FaceModelRepository) GetByLabel(db *gorm.DB, label uint) *model.FaceModel {
	faceModel := new(model.FaceModel)
	result := db.Where("id = ?", label).First(faceModel)
	if result.RowsAffected == 1 {
		return faceModel
	}

	return nil
}

func (faceModelRepository FaceModelRepository) GetDeviceFaceModels(db *gorm.DB, deviceId uint, page uint, pageSize uint) *repository.PageModel {
	faceModels := make([]model.FaceModel, 0)
	pageData := &repository.PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &faceModels,
	}

	totalQuery := db.Model(&model.FaceModel{}).Where("device_id = ?", deviceId).Where("status != ?", model.FACE_MODEL_STATUS_DISABLE)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&faceModels)
		pageData.Total = uint64(total)
	}

	return pageData
}
