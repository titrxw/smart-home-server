package main

import (
	global "github.com/titrxw/go-framework/src/Global"
	app "github.com/titrxw/smart-home-server/app"
	"github.com/titrxw/smart-home-server/server"
	"strconv"
)

func main() {
	app := app.NewApp()
	app.Bootstrap()

	server.RegisterHttpServer(app)

	global.FHttpServer.Start(app.Config.Server.Host + ":" + strconv.Itoa(app.Config.Server.Port))
	//app.Console.Run()
}
