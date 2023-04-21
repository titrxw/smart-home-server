package middleware

import (
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/user"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Oauth struct {
	middleware.Abstract
	user.Oauth
}

func (m Oauth) Process(ctx *gin.Context) {
	userId := m.GetUserIdFromSession(ctx)
	if userId <= 0 {
		m.JsonResponseWithError(ctx, exception.NewResponseError("未登录"), http.StatusForbidden)
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
