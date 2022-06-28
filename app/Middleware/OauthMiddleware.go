package middleware

import (
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/titrxw/smart-home-server/app/Controller/Frontend/User"
)

type OauthMiddleware struct {
	middleware.MiddlewareAbstract
	user.UserOauth
}

func (this OauthMiddleware) Process(ctx *gin.Context) {
	userId := this.GetUserIdFromSession(ctx)
	if userId <= 0 {
		this.JsonResponseWithError(ctx, "未登录", http.StatusForbidden)
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
