package repository

import (
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type AppProxyRepository struct {
	RepositoryAbstract
}

func (appProxyRepository AppProxyRepository) AddAppProxy(db *gorm.DB, app *model.App, componentApp *model.App) *model.AppProxy {
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

func (appProxyRepository AppProxyRepository) GetComponentAppProxyAppId(db *gorm.DB, componentApp *model.App) string {
	appProxyModel := new(model.AppProxy)
	result := db.Where("component_app_id = ?", componentApp.AppId).First(appProxyModel)
	if result.RowsAffected == 1 {
		return appProxyModel.AppId
	}

	return ""
}

func (appProxyRepository AppProxyRepository) ClearComponentAppProxy(db *gorm.DB, componentApp *model.App) error {
	result := db.Where("component_app_id = ?", componentApp.AppId).Delete(&model.AppProxy{})
	return result.Error
}
