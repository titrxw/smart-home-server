package logic

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/internal/repository"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	"gorm.io/gorm"
)

type App struct {
	logic.Abstract
}

func (l *App) CreateApp(appType uint8) *model.App {
	app := repository.Repository.App.CreateApp(l.GetDefaultDb(), appType)
	if app == nil {
		return nil
	}

	return app
}

func (l *App) GetAppByAppId(appid string) *model.App {
	app := repository.Repository.App.GetByAppId(l.GetDefaultDb(), appid)
	if app == nil {
		return nil
	}

	return app
}

func (l *App) GetOpenAppByAppId(appid string) *model.App {
	app := repository.Repository.App.GetByAppIdAndType(l.GetDefaultDb(), appid, model.OpenAppType)
	if app == nil {
		return nil
	}

	return app
}

func (l *App) GetDeviceAppByAppId(appid string) *model.App {
	app := repository.Repository.App.GetByAppIdAndType(l.GetDefaultDb(), appid, model.DeviceAppType)
	if app == nil {
		return nil
	}

	return app
}

func (l *App) addAppProxy(app *model.App, componentApp *model.App) error {
	return l.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		err := repository.Repository.AppProxy.ClearComponentAppProxy(tx, componentApp)
		if err != nil {
			return err
		}
		appProxy := repository.Repository.AppProxy.AddAppProxy(tx, app, componentApp)
		if appProxy == nil {
			return exception.NewResponseError("设备绑定失败")
		}

		return nil
	})
}
