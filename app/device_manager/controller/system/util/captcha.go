package util

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
	"github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type Captcha struct {
	controller.Abstract
}

func (c Captcha) Captcha(ctx *gin.Context) {
	captchaId, b64s, err := logic.Logic.Captcha.GenerateCaptcha(ctx)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	err = server.GHttpServer.Session.Set(ctx, "captcha", captchaId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	ctx.Writer.WriteString(b64s)
}
