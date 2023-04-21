package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/url"
)

type AppCheck struct {
	middleware.Abstract
}

func (m AppCheck) Process(ctx *gin.Context) {
	var params map[string]string
	var form url.Values
	if ctx.ContentType() == "application/x-www-form-urlencoded" {
		form = ctx.Request.PostForm
	}
	if ctx.ContentType() == "multipart/form-data" {
		muForm, err := ctx.MultipartForm()
		if err != nil {
			m.JsonResponseWithServerError(ctx, err)
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
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("appid缺失"))
		return
	}
	psign, ok := params["sign"]
	if !ok {
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("sign缺失"))
		return
	}
	_, ok = params["nonce"]
	if !ok {
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("nonce缺失"))
		return
	}
	_, ok = params["timestamp"]
	if !ok {
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("timestamp缺失"))
		return
	}

	app := logic.Logic.App.GetAppByAppId(pappid)
	if app == nil {
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("appid错误"))
		return
	}

	sign := logic.Logic.App.GetSign(app, params)
	if psign != sign {
		m.JsonResponseWithServerError(ctx, exception.NewResponseError("签名错误"))
		return
	}

	ctx.Set("app", app)
	ctx.Next()
}
