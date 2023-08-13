package repository

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/repository"
	"gorm.io/gorm"
)

type AppProxy struct {
	repository.Abstract
}

func (r AppProxy) AddAppProxy(db *gorm.DB, app *model.App, componentApp *model.App) *model.AppProxy {
	appProxyModel := new(model.AppProxy)
	result := db.Where("app_id = ?", app.AppId).Where("component_app_id = ?", componentApp.AppId).First(appProxyModel)
	if result.RowsAffected == 1 {
		return appProxyModel
	}

	appProxyModel.AppId = app.AppId
	appProxyModel.ComponentAppId = componentApp.AppId
	result = db.Create(appProxyModel)
	if result.RowsAffected == 1 {
		return appProxyModel
	}

	return nil
}

func (r AppProxy) GetComponentAppProxyAppId(db *gorm.DB, componentApp *model.App) string {
	appProxyModel := new(model.AppProxy)
	result := db.Where("component_app_id = ?", componentApp.AppId).First(appProxyModel)
	if result.RowsAffected == 1 {
		return appProxyModel.AppId
	}

	return ""
}

func (r AppProxy) ClearComponentAppProxy(db *gorm.DB, componentApp *model.App) error {
	result := db.Where("component_app_id = ?", componentApp.AppId).Delete(&model.AppProxy{})
	return result.Error
}
