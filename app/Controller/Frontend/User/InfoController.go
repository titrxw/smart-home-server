package user

import (
	"github.com/gin-gonic/gin"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type InfoController struct {
	frontend.ControllerAbstract
}

func (this InfoController) Info(ctx *gin.Context) {
	user, err := logic.Logic.UserLogic.GetUserById(ctx, this.GetUserId(ctx))
	if err != nil {
		this.JsonResponseWithServerError(ctx, err)
		return
	}
	this.JsonResponseWithoutError(ctx, user)
}
