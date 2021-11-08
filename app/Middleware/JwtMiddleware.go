package middleware

import "github.com/gin-gonic/gin"

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
