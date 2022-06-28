package frontend

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type ControllerAbstract struct {
	base.ControllerAbstract
}

func (this ControllerAbstract) GetUserId(ctx *gin.Context) model.UID {
	userId, _ := ctx.MustGet("user_id").(model.UID)
	return userId
}
