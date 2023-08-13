package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/http"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	pkgmiddleware "github.com/titrxw/smart-home-server/app/pkg/middleware"
)

type AppCheck struct {
	pkgmiddleware.SignCheck
}

func (m AppCheck) ValidateSign(ctx *gin.Context, params map[string]string) error {
	app := logic.Logic.App.GetAppByAppId(params["appid"])
	if app == nil {
		return exception.NewResponseError("appid错误")
	}

	rsign := params["sign"]
	sign := http.GetSign(app.AppSecret, params)
	if rsign != sign {
		return exception.NewResponseError("签名错误")
	}

	ctx.Set("app", app)

	return nil
}

func (m AppCheck) Process(ctx *gin.Context) {
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
