package main

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server"
	"github.com/titrxw/smart-home-server/router"
	"strconv"
)

func main() {
	app := app.NewApp()
	app.Bootstrap()

	gin.SetMode(app.Config.App.Env)
	server := gin.Default()
	router.Register(server)
	server.Run(app.Config.Server.Host + ":" + strconv.Itoa(app.Config.Server.Port))
}
