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

func (this *App) Bootstrap() {
	this.App.Bootstrap()
	this.InitConfig(&this.Config)

	this.App.HandlerExceptions.SetExceptionHandler(new(exception.ExceptionHandler))

	this.ProviderManager.MakeProvider(new(provider.ServiceProvider)).Register(this.Config)
	this.ProviderManager.MakeProvider(new(provider.ValidatorProvider)).Register(this.Config)
	this.ProviderManager.MakeProvider(new(provider.DeviceProvider)).Register(this.Config)
}
