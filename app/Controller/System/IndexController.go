package system

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
)

type IndexController struct {
	base.ControllerAbstract
}

func (indexController *IndexController) Index(ctx *gin.Context) {
	indexController.JsonResponseWithoutError(ctx, "")
}
