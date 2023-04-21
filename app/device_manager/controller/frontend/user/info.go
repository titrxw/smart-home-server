package user

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
)

type Info struct {
	frontend.Abstract
}

func (c Info) Info(ctx *gin.Context) {
	user, err := logic.Logic.User.GetUserById(ctx.Request.Context(), c.GetUserId(ctx))
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, user)
}
