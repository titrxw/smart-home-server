package app

import (
	app "github.com/titrxw/go-framework/src/App"
	global "github.com/titrxw/go-framework/src/Global"
	exception "github.com/titrxw/smart-home-server/app/Handler/Exception"
	provider "github.com/titrxw/smart-home-server/app/Provider"
	"github.com/titrxw/smart-home-server/config"
)

var GApp *App

type App struct {
	*app.App
	Config *config.Config
}

func NewApp() *App {
	GApp = &App{
		App: app.NewApp(),
	}
	global.FApp = GApp.App
	return GApp
}

func (app *App) Bootstrap() {
	app.App.Bootstrap()
	app.InitConfig(&app.Config)

	app.App.HandlerExceptions.SetExceptionHandler(new(exception.ExceptionHandler))

	app.ProviderManager.MakeProvider(new(provider.ServiceProvider)).Register(app.Config)
	app.ProviderManager.MakeProvider(new(provider.ValidatorProvider)).Register(app.Config)
	app.ProviderManager.MakeProvider(new(provider.DeviceProvider)).Register(app.Config)
}
