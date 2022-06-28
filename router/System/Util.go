package system

import (
	"github.com/gin-gonic/gin"
	util "github.com/titrxw/smart-home-server/app/Controller/System/Util"
)

type Util struct {
}

func (this *Util) registerRoute(router *gin.RouterGroup) {
	v2 := router.Group("/util")
	{
		v2.GET("/captcha", new(util.CaptchaController).Captcha)
	}
}
