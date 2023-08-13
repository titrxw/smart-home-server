package main

import (
	"bytes"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	deviceManagerApp "github.com/titrxw/smart-home-server/app/device_manager"
	"github.com/titrxw/smart-home-server/app/devices/face_identify"
	"github.com/titrxw/smart-home-server/app/devices/mqtt_zigbee_gateway"
	"github.com/titrxw/smart-home-server/app/mqtt"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"github.com/we7coreteam/w7-rangine-go/src/core/helper"
	"github.com/we7coreteam/w7-rangine-go/src/http"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"github.com/we7coreteam/w7-rangine-go/src/http/session"
)

//go:embed config.yaml
var ConfigFileContent []byte

func main() {
	newApp := app.NewApp(app.Option{
		DefaultConfigLoader: func(config *viper.Viper) {
			config.SetConfigType("yaml")
			err := config.MergeConfig(bytes.NewReader(helper.ParseConfigContentEnv(ConfigFileContent)))
			if err != nil {
				panic(err)
			}
		},
	})

	httpServer := new(http.Provider).Register(newApp.GetConfig(), newApp.GetConsole(), newApp.GetServerManager()).Export()
	mqttServer := new(mqtt.Provider).Register(newApp.GetConfig(), newApp.GetServerManager()).Export()

	httpServer.Use(middleware.GetPanicHandlerMiddleware())
	httpServer.Use(middleware.GetSessionMiddleware(newApp.GetConfig(), session.GetGormStore, []byte("secret")))

	httpServer.RegisterRouters(func(engine *gin.Engine) {
		engine.Static("/static", "./public/static")
		engine.Static("/img", "./public/upload/img")
	})

	deviceManagerApp.Provider{}.Register(httpServer, newApp.GetConsole())
	face_identify.Provider{}.Register(httpServer, newApp.GetConfig())
	mqtt_zigbee_gateway.Provider{}.Register(mqttServer)

	newApp.RunConsole()
}
