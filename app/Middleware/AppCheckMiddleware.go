package middleware

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/titrxw/go-framework/src/Http/Middleware"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	http "github.com/titrxw/smart-home-server/app/Http"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"net/url"
)

type AppCheckMiddleware struct {
	middleware.MiddlewareAbstract
	http.ResponseFormat
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
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewArgsError("appid缺失"))
		return
	}
	psign, ok := params["sign"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewArgsError("sign缺失"))
		return
	}
	_, ok = params["nonce"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewArgsError("nonce缺失"))
		return
	}
	_, ok = params["timestamp"]
	if !ok {
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewArgsError("timestamp缺失"))
		return
	}

	app := logic.Logic.AppLogic.GetAppByAppId(pappid)
	if app == nil {
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewLogicError("appid错误"))
		return
	}

	sign := logic.Logic.AppLogic.GetSign(app, params)
	if psign != sign {
		appCheckMiddleware.JsonResponseWithServerError(ctx, exception.NewLogicError("签名错误"))
		return
	}

	ctx.Set("app", app)
	ctx.Next()
}
