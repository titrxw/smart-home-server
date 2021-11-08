package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/titrxw/smart-home-server"
	"github.com/titrxw/smart-home-server/router"
)

func main() {
	app := app.NewApp()
	app.Bootstrap()

	gin.SetMode(app.Config.App.Env)
	server := gin.Default()
	router.GRouter.Register(server)
	server.Run(app.Config.Server.Host + ":" + strconv.Itoa(app.Config.Server.Port))
}
