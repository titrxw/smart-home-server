package middleware

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"net/url"
)

type AppCheckMiddleware struct {
	middleware.MiddlewareAbstract
}

func (appCheckMiddleware *AppCheckMiddleware) Process(ctx *gin.Context) {
	var params map[string]string
	var form url.Values
	if ctx.ContentType() == "application/x-www-form-urlencoded" {
		form = ctx.Request.PostForm
	}
	if ctx.ContentType() == "multipart/form-data" {
		muForm, err := ctx.MultipartForm()
		if err != nil {
			appCheckMiddleware.JsonResponseWithServerError(ctx, err)
			return
		}
		form = muForm.Value
	}

	query := ctx.Request.URL.Query()
	params = make(map[string]string, len(query)+len(form))

	for k := range query {
		params[k] = ctx.Query(k)
	}
	for k := range form {
		params[k] = ctx.PostForm(k)
	}

	pappid, ok := params["appid"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "appid缺失")
		return
	}
	psign, ok := params["sign"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "sign缺失")
		return
	}
	_, ok = params["nonce"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "nonce缺失")
		return
	}
	_, ok = params["timestamp"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "timestamp缺失")
		return
	}

	app := logic.Logic.AppLogic.GetAppByAppId(pappid)
	if app == nil {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "appid错误")
		return
	}

	sign := logic.Logic.AppLogic.GetSign(app, params)
	if psign != sign {
		appCheckMiddleware.JsonResponseWithServerError(ctx, "签名错误")
		return
	}

	ctx.Set("app", app)
	ctx.Next()
}
