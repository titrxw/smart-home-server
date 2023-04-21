package main

import (
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/common"
	deviceManagerApp "github.com/titrxw/smart-home-server/app/device_manager"
	"github.com/titrxw/smart-home-server/app/devices/face_identify"
	"github.com/titrxw/smart-home-server/app/devices/light"
	"github.com/titrxw/smart-home-server/app/devices/mqtt_zigbee_gateway"
	"github.com/titrxw/smart-home-server/app/mqtt"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"github.com/we7coreteam/w7-rangine-go/src/http/server"
)

func main() {
	newApp := app.NewApp()

	server.GetSession().SetStorageResolver(func() scs.Store {
		sessionDb := newApp.GetConfig().GetString("session.db")
		if sessionDb == "" {
			sessionDb = "default"
		}
		db, err := facade.GetDbFactory().Channel(sessionDb)
		if err != nil {
			panic(err)
		}
		HSQLDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		return mysqlstore.New(HSQLDB)
	})
	server.Use(middleware.PanicHandlerMiddleware{}.GetProcess())
	server.Use(middleware.NewSessionMiddleware(server.GetSession()).Process)

	server.RegisterRouters(func(engine *gin.Engine) {
		engine.Static("/static", "./public/static")
		engine.Static("/img", "./public/upload/img")
	})

	newApp.GetProviderManager().RegisterProvider(new(common.Provider)).Register()
	newApp.GetProviderManager().RegisterProvider(new(deviceManagerApp.Provider)).Register()
	newApp.GetProviderManager().RegisterProvider(new(mqtt.Provider)).Register()
	newApp.GetProviderManager().RegisterProvider(new(face_identify.Provider)).Register()
	newApp.GetProviderManager().RegisterProvider(new(light.Provider)).Register()
	newApp.GetProviderManager().RegisterProvider(new(mqtt_zigbee_gateway.Provider)).Register()

	newApp.RunConsole()
}
