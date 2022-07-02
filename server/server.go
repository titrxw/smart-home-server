package server

import (
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	global "github.com/titrxw/go-framework/src/Global"
	route "github.com/titrxw/go-framework/src/Http/Command/Route"
	server2 "github.com/titrxw/go-framework/src/Http/Command/Server"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	server "github.com/titrxw/go-framework/src/Http/Server"
	session2 "github.com/titrxw/go-framework/src/Http/Session"
	"github.com/titrxw/smart-home-server/app"
	"github.com/titrxw/smart-home-server/router"
)

func RegisterHttpServer(app *app.App) {
	global.FHttpServer = server.NewHttpSerer(app.App)

	global.FHttpServer.Session = session2.NewSession(app.App.Config.Session, app.App.Config.Cookie)
	global.FHttpServer.Session.SetStorageResolver(func() scs.Store {
		sessionDb := app.Config.Session.DbConnection
		if sessionDb == "" {
			sessionDb = "default"
		}
		db, err := app.App.DbFactory.Channel(sessionDb).DB()
		if err != nil {
			panic(err)
		}

		return mysqlstore.New(db)
	})

	global.FHttpServer.GinEngine.Use(middleware.ExceptionMiddleware{HandlerExceptions: app.App.HandlerExceptions}.Process)
	global.FHttpServer.GinEngine.Use(middleware.NewSessionMiddleware(global.FHttpServer.Session).Process)

	global.FHttpServer.RegisterRouters(func(engine *gin.Engine) {
		router.GRouter.Register(engine)
	})

	app.Console.RegisterCommand(&server2.StartCommand{
		Server: global.FHttpServer,
	})
	app.Console.RegisterCommand(&route.ListCommand{
		GinEngine: global.FHttpServer.GinEngine,
	})
}
