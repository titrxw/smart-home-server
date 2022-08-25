package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/titrxw/go-framework/src/Http/Response"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	"net/http"
)

type ResponseFormat struct {
}

func (responseFormat ResponseFormat) JsonResponseWithServerError(ctx *gin.Context, err interface{}) {
	responseFormat.JsonResponseWithError(ctx, err, http.StatusInternalServerError)
}

func (responseFormat ResponseFormat) JsonResponseWithError(ctx *gin.Context, err interface{}, statusCode int) {
	response := Response.Response{}
	switch err.(type) {
	case exception.ArgsError:
		response.JsonResponse(ctx, "", err.(*exception.ArgsError).Error(), statusCode)
	case exception.LogicError:
		response.JsonResponse(ctx, "", err.(*exception.LogicError).Error(), statusCode)
	case error:
		if gin.IsDebugging() {
			response.JsonResponse(ctx, "", err.(error).Error(), statusCode)
		} else {
			response.JsonResponse(ctx, "", errors.New("系统内部错误"), statusCode)
		}
	default:
		response.JsonResponse(ctx, "", err, statusCode)
	}
}
