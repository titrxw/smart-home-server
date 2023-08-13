package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/http"
	"strconv"
)

type UserId struct {
	middleware.Abstract
}

func (m UserId) Process(ctx *gin.Context) {
	result := ctx.PostForm("user_id")
	if result == "" {
		m.JsonResponseWithError(ctx, exception.NewResponseError("user_id参数错误"), http.StatusForbidden)
		ctx.Abort()
		return
	}

	userId, err := strconv.Atoi(result)
	if err != nil {
		m.JsonResponseWithError(ctx, exception.NewResponseError("user_id参数错误"), http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
