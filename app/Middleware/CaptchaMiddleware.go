package middleware

import (
	global "github.com/titrxw/go-framework/src/Global"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	captcha "github.com/titrxw/smart-home-server/app/Utils/Captcha"
)

type CaptchaMiddleware struct {
	middleware.MiddlewareAbstract
}

func (this *CaptchaMiddleware) Process(ctx *gin.Context) {
	captchaValue := ctx.PostForm("captcha")
	if captchaValue == "" {
		this.JsonResponseWithError(ctx, "验证码参数错误", http.StatusForbidden)
		return
	}

	captchaId, ok := global.FHttpServer.Session.Get(ctx, "captcha").(string)
	if !ok {
		this.JsonResponseWithError(ctx, "验证码过期", http.StatusForbidden)
		return
	}

	if !captcha.NewDefaultRedisStore(global.FApp.RedisFactory.Channel("default"), ctx).Verify(captchaId, captchaValue, true) {
		this.JsonResponseWithError(ctx, "验证码错误", http.StatusForbidden)
		return
	}

	ctx.Next()
}
