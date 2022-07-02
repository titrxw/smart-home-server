package main

import (
	"strconv"

	global "github.com/titrxw/go-framework/src/Global"
	app "github.com/titrxw/smart-home-server/app"
	"github.com/titrxw/smart-home-server/mqtt"
	"github.com/titrxw/smart-home-server/server"
)

func main() {
	app := app.NewApp()
	app.Bootstrap()

	server.RegisterHttpServer(app)
	mqtt.RegisterSubscribe(app)

	global.FHttpServer.Start(app.Config.Server.Host + ":" + strconv.Itoa(app.Config.Server.Port))
	//app.Console.Run()
}
