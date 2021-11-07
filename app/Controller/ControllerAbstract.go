package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerAbstract struct {
}

func (this *ControllerAbstract) jsonSuccessResponse(ctx *gin.Context) {
	this.jsonResponseWithoutError(ctx, "success")
}

func (this *ControllerAbstract) jsonResponseWithoutError(ctx *gin.Context, data interface{}) {
	this.jsonResponse(ctx, data, "", http.StatusOK)
}

func (this *ControllerAbstract) jsonResponse(ctx *gin.Context, data interface{}, error string, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"data":  data,
		"code":  http.StatusOK,
		"error": error,
	})
}
