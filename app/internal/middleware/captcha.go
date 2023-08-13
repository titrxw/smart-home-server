package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
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
		ctx.Abort()
		return
	}

	captchaId, ok := sessions.Default(ctx).Get("captcha").(string)
	if !ok {
		m.JsonResponseWithError(ctx, exception.NewResponseError("验证码过期"), http.StatusForbidden)
		ctx.Abort()
		return
	}

	if !logic.Logic.Captcha.ValidateCaptcha(ctx, captchaId, captchaValue) {
		m.JsonResponseWithError(ctx, exception.NewResponseError("验证码错误"), http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Next()
}
