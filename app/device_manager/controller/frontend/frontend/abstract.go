package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
)

type Abstract struct {
	controller.Abstract
}

func (c Abstract) GetUserId(ctx *gin.Context) model.UID {
	userId, _ := ctx.MustGet("user_id").(model.UID)
	return userId
}
