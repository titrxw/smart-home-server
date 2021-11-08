package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerAbstract struct {
}

func (this *ControllerAbstract) JsonSuccessResponse(ctx *gin.Context) {
	this.JsonResponseWithoutError(ctx, "success")
}

func (this *ControllerAbstract) JsonResponseWithoutError(ctx *gin.Context, data interface{}) {
	this.JsonResponse(ctx, data, "", http.StatusOK)
}

func (this *ControllerAbstract) JsonResponse(ctx *gin.Context, data interface{}, error string, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"data":  data,
		"code":  http.StatusOK,
		"error": error,
	})
}
