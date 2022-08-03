package user

import (
	"github.com/gin-gonic/gin"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type InfoController struct {
	frontend.ControllerAbstract
}

func (infoController InfoController) Info(ctx *gin.Context) {
	user, err := logic.Logic.UserLogic.GetUserById(ctx, infoController.GetUserId(ctx))
	if err != nil {
		infoController.JsonResponseWithServerError(ctx, err)
		return
	}

	infoController.JsonResponseWithoutError(ctx, user)
}
