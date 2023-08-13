package repository

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/repository"
	"strings"

	"gorm.io/gorm"
)

type App struct {
	repository.Abstract
}

func (r App) CreateApp(db *gorm.DB, appType uint8) *model.App {
	appId := r.getUniqueAppId(db, "sd", appType)
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

func (r App) GetByAppIdAndType(db *gorm.DB, appId string, appType uint8) *model.App {
	appModel := new(model.App)
	result := db.Where("app_id = ?", appId).Where("app_type = ?", appType).First(appModel)
	if result.RowsAffected == 1 {
		return appModel
	}

	return nil
}

func (r App) getUniqueAppId(db *gorm.DB, prefix string, appType uint8) string {
	appId := prefix + helper.RandomStr(22)
	if r.GetByAppIdAndType(db, appId, appType) != nil {
		return r.getUniqueAppId(db, prefix, appType)
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
