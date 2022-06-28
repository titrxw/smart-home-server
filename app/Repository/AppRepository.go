package repository

import (
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
	"strings"
)

type AppRepository struct {
	RepositoryAbstract
}

func (this AppRepository) CreateDeviceApp(db *gorm.DB) *model.App {
	appId := this.getUniqueAppId(db, "sd")
	return this.make(db, appId, strings.Replace(helper.UUid(), "-", "", -1), model.DEVICE_APP_TYPE)
}

func (this AppRepository) GetById(db *gorm.DB, id uint) *model.App {
	appModel := new(model.App)
	result := db.Where("id = ?", id).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (this AppRepository) GetByAppId(db *gorm.DB, appId string) *model.App {
	appModel := new(model.App)
	result := db.Where("app_id = ?", appId).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (this AppRepository) getUniqueAppId(db *gorm.DB, prefix string) string {
	appId := prefix + helper.RandomStr(22)
	if this.GetByAppId(db, appId) != nil {
		return this.getUniqueAppId(db, prefix)
	}

	return appId
}

func (this AppRepository) make(db *gorm.DB, appid string, appSecret string, appType uint8) *model.App {
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
