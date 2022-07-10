package middleware

import (
	"net/http"

	middleware "github.com/titrxw/go-framework/src/Http/Middleware"

	"github.com/gin-gonic/gin"
	user "github.com/titrxw/smart-home-server/app/Controller/Frontend/User"
)

type OauthMiddleware struct {
	middleware.MiddlewareAbstract
	user.UserOauth
}

func (oauthMiddleware OauthMiddleware) Process(ctx *gin.Context) {
	userId := oauthMiddleware.GetUserIdFromSession(ctx)
	if userId <= 0 {
		oauthMiddleware.JsonResponseWithError(ctx, "未登录", http.StatusForbidden)
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}