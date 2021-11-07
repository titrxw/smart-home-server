package controller

import "github.com/gin-gonic/gin"

type IndexController struct {
	ControllerAbstract
}

func (this *IndexController) Index(ctx *gin.Context) {
	this.jsonSuccessResponse(ctx)
}
