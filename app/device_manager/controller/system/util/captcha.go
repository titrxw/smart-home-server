package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
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

	sessions.Default(ctx).Set("captcha", captchaId)
	err = sessions.Default(ctx).Save()
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	ctx.Writer.WriteString(b64s)
}
