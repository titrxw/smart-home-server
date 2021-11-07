package app

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	appconfig "github.com/titrxw/smart-home-server/config"
)

var GApp *App

type App struct {
	Config *appconfig.Config
}

func NewApp() *App {
	GApp = &App{}
	return GApp
}

func (this *App) registerConfig() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("./app_config.yaml")
	if err != nil {
		panic(err)
	}
	err = config.BindStruct("", &this.Config)
	if err != nil {
		panic(err)
	}
}

func (this *App) Bootstrap() {
	this.registerConfig()
}
