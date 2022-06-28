package system

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
)

type IndexController struct {
	base.ControllerAbstract
}

func (this *IndexController) Index(ctx *gin.Context) {
	this.JsonResponseWithoutError(ctx, "")
}
