package frontend

import (
	"github.com/gin-gonic/gin"
	user "github.com/titrxw/smart-home-server/app/Controller/Frontend/User"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
)

type User struct {
}

func (userr *User) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/oauth")
	{
		v2.POST("/register-email", new(middleware.CaptchaMiddleware).Process, new(user.OauthController).SendRegisterEmailCode)
		v2.POST("/register", new(user.OauthController).Register)
		v2.POST("/login", new(middleware.CaptchaMiddleware).Process, new(user.OauthController).Login)
		v2.GET("/logout", new(middleware.OauthMiddleware).Process, new(user.OauthController).Logout)
		v2.GET("/info", new(middleware.OauthMiddleware).Process, new(user.InfoController).Info)
	}
}
