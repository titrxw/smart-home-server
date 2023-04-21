package repository

import (
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"strings"

	"gorm.io/gorm"
)

type App struct {
	Abstract
}

func (r App) CreateDeviceApp(db *gorm.DB, appType uint8) *model.App {
	appId := r.getUniqueAppId(db, "sd")
	return r.make(db, appId, strings.Replace(helper.UUid(), "-", "", -1), appType)
}

func (r App) GetById(db *gorm.DB, id uint) *model.App {
	appModel := new(model.App)
	result := db.Where("id = ?", id).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (r App) GetByAppId(db *gorm.DB, appId string) *model.App {
	appModel := new(model.App)
	result := db.Where("app_id = ?", appId).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (r App) getUniqueAppId(db *gorm.DB, prefix string) string {
	appId := prefix + helper.RandomStr(22)
	if r.GetByAppId(db, appId) != nil {
		return r.getUniqueAppId(db, prefix)
	}

	return appId
}

func (r App) make(db *gorm.DB, appid string, appSecret string, appType uint8) *model.App {
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
