package repository

import (
	"strings"

	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type AppRepository struct {
	RepositoryAbstract
}

func (appRepository AppRepository) CreateDeviceApp(db *gorm.DB) *model.App {
	appId := appRepository.getUniqueAppId(db, "sd")
	return appRepository.make(db, appId, strings.Replace(helper.UUid(), "-", "", -1), model.DEVICE_APP_TYPE)
}

func (appRepository AppRepository) GetById(db *gorm.DB, id uint) *model.App {
	appModel := new(model.App)
	result := db.Where("id = ?", id).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (appRepository AppRepository) GetByAppId(db *gorm.DB, appId string) *model.App {
	appModel := new(model.App)
	result := db.Where("app_id = ?", appId).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (appRepository AppRepository) getUniqueAppId(db *gorm.DB, prefix string) string {
	appId := prefix + helper.RandomStr(22)
	if appRepository.GetByAppId(db, appId) != nil {
		return appRepository.getUniqueAppId(db, prefix)
	}

	return appId
}

func (appRepository AppRepository) make(db *gorm.DB, appid string, appSecret string, appType uint8) *model.App {
	appModel := &model.App{
		AppId:     appid,
		AppSecret: appSecret,
		AppType:   appType,
	}

	result := db.Create(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}
