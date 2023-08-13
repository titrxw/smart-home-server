package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/http"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"net/url"
)

type SignCheck struct {
	middleware.Abstract

	AppId     string
	AppSecret string
}

func (m SignCheck) ValidateParams(ctx *gin.Context) (map[string]string, error) {
	var params map[string]string
	var form url.Values
	if ctx.ContentType() == "application/x-www-form-urlencoded" {
		err := ctx.Request.ParseForm()
		if err != nil {
			return nil, err
		}
		form = ctx.Request.PostForm
	}
	if ctx.ContentType() == "multipart/form-data" {
		muForm, err := ctx.MultipartForm()
		if err != nil {
			return nil, err
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

	_, ok := params["appid"]
	if !ok {
		return nil, exception.NewResponseError("appid缺失")
	}
	_, ok = params["sign"]
	if !ok {
		return nil, exception.NewResponseError("sign缺失")
	}
	_, ok = params["nonce"]
	if !ok {
		return nil, exception.NewResponseError("nonce缺失")
	}
	_, ok = params["timestamp"]
	if !ok {
		return nil, exception.NewResponseError("timestamp缺失")
	}

	_, ok = params["body"]
	if !ok {
		return nil, exception.NewResponseError("body缺失")
	}

	return params, nil
}

func (m SignCheck) ValidateSign(ctx *gin.Context, params map[string]string) error {
	if m.AppId != params["appid"] {
		return exception.NewResponseError("appid错误")
	}

	psign := params["sign"]
	sign := http.GetSign(m.AppSecret, params)
	if psign != sign {
		return exception.NewResponseError("签名错误")
	}

	return nil
}

func (m SignCheck) Process(ctx *gin.Context) {
	params, err := m.ValidateParams(ctx)
	if err != nil {
		m.JsonResponseWithServerError(ctx, err)
		ctx.Abort()
		return
	}

	err = m.ValidateSign(ctx, params)
	if err != nil {
		m.JsonResponseWithServerError(ctx, err)
		ctx.Abort()
		return
	}

	ctx.Next()
}
