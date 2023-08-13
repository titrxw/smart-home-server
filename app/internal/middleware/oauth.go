package middleware

import (
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Oauth struct {
	middleware.Abstract
	UserOauth
}

func (m Oauth) Process(ctx *gin.Context) {
	userId := m.GetUserIdFromSession(ctx)
	if userId <= 0 {
		m.JsonResponseWithError(ctx, exception.NewResponseError("未登录"), http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
