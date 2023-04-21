package middleware

import (
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"github.com/we7coreteam/w7-rangine-go/src/http/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Captcha struct {
	middleware.Abstract
}

func (m Captcha) Process(ctx *gin.Context) {
	captchaValue := ctx.PostForm("captcha")
	if captchaValue == "" {
		m.JsonResponseWithError(ctx, exception.NewResponseError("验证码参数错误"), http.StatusForbidden)
		return
	}

	captchaId, ok := server.GHttpServer.Session.Get(ctx, "captcha").(string)
	if !ok {
		m.JsonResponseWithError(ctx, exception.NewResponseError("验证码过期"), http.StatusForbidden)
		return
	}

	if !logic.Logic.Captcha.ValidateCaptcha(ctx, captchaId, captchaValue) {
		m.JsonResponseWithError(ctx, exception.NewResponseError("验证码错误"), http.StatusForbidden)
		return
	}

	ctx.Next()
}
