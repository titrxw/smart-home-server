package middleware

import "github.com/gin-gonic/gin"

func JwtMiddleware(ctx *gin.Context) {
	ctx.Next()
}
