package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	global "github.com/titrxw/go-framework/src/Global"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	captcha "github.com/titrxw/smart-home-server/app/Utils/Captcha"
)

type CaptchaController struct {
	base.ControllerAbstract
}

func (captchaController CaptchaController) Captcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captcha.NewDefaultRedisStore(global.FApp.RedisFactory.Channel("default"), ctx))
	captchaId, b64s, err := cp.Generate()
	if err != nil {
		captchaController.JsonResponseWithServerError(ctx, err)
		return
	}

	err = global.FHttpServer.Session.Set(ctx, "captcha", captchaId)
	if err != nil {
		captchaController.JsonResponseWithServerError(ctx, err)
		return
	}

	ctx.Writer.WriteString(b64s)
}
