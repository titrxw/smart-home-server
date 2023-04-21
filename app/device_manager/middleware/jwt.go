package middleware

import (
	"github.com/titrxw/smart-home-server/app/common/jwt"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Jwt struct {
	middleware.Abstract
}

func (m Jwt) Process(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		m.JsonResponseWithError(ctx, exception.NewResponseError("请求头中auth为空"), http.StatusBadRequest)
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	token, err := jwt.GetJwtService().ParseToken(authHeader)
	if err != nil {
		m.JsonResponseWithError(ctx, exception.NewResponseError("无效的Token"), http.StatusForbidden)
		return
	}
	err = jwt.GetJwtService().ValidateToken(token)
	if err != nil {
		m.JsonResponseWithError(ctx, exception.NewResponseError("无效的Token"), http.StatusForbidden)
		return
	}

	ctx.Set("user", token.Claims.(jwt.Claims).Payload)
	ctx.Next()
}
