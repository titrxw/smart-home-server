package middleware

import (
	exception "github.com/titrxw/smart-home-server/app/Exception"
	http2 "github.com/titrxw/smart-home-server/app/Http"
	"net/http"

	global "github.com/titrxw/go-framework/src/Global"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"

	"github.com/gin-gonic/gin"
	captcha "github.com/titrxw/smart-home-server/app/Utils/Captcha"
)

type CaptchaMiddleware struct {
	middleware.MiddlewareAbstract
	http2.ResponseFormat
}

func (captchaMiddleware *CaptchaMiddleware) Process(ctx *gin.Context) {
	captchaValue := ctx.PostForm("captcha")
	if captchaValue == "" {
		captchaMiddleware.JsonResponseWithError(ctx, exception.NewArgsError("验证码参数错误"), http.StatusForbidden)
		return
	}

	captchaId, ok := global.FHttpServer.Session.Get(ctx, "captcha").(string)
	if !ok {
		captchaMiddleware.JsonResponseWithError(ctx, exception.NewLogicError("验证码过期"), http.StatusForbidden)
		return
	}

	if !captcha.NewDefaultRedisStore(global.FApp.RedisFactory.Channel("default"), ctx).Verify(captchaId, captchaValue, true) {
		captchaMiddleware.JsonResponseWithError(ctx, exception.NewLogicError("验证码错误"), http.StatusForbidden)
		return
	}

	ctx.Next()
}
