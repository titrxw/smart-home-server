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
