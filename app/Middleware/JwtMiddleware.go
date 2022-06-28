package middleware

import (
	"github.com/gin-gonic/gin"
	global "github.com/titrxw/go-framework/src/Global"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	jwt "github.com/titrxw/smart-home-server/app/Service/Jwt"
	"net/http"
)

type JwtMiddleware struct {
	middleware.MiddlewareAbstract
}

func (this JwtMiddleware) Process(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		this.JsonResponseWithError(ctx, "请求头中auth为空", http.StatusBadRequest)
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	token, err := jwt.GetJwtService(global.FApp.Container).ParseToken(authHeader)
	if err != nil {
		this.JsonResponseWithError(ctx, "无效的Token", http.StatusForbidden)
		return
	}
	err = jwt.GetJwtService(global.FApp.Container).ValidateToken(token)
	if err != nil {
		this.JsonResponseWithError(ctx, "无效的Token", http.StatusForbidden)
		return
	}

	ctx.Set("user", token.Claims.(jwt.Claims).Payload)
	ctx.Next()
}
