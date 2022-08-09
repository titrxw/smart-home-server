package system

import (
	"github.com/gin-gonic/gin"
	util "github.com/titrxw/smart-home-server/app/Controller/System/Util"
	middleware "github.com/titrxw/smart-home-server/app/Middleware"
)

type Util struct {
}

func (urilr *Util) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/util")
	{
		v2.GET("/captcha", new(util.CaptchaController).Captcha)
		v2.POST("/attach/upload/image", new(middleware.AppCheckMiddleware).Process, new(util.UploadController).UploadImage)
	}
}
