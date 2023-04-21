package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/device_manager/repository"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type App struct {
	Abstract
}

func (l *App) GetAppByAppId(appid string) *model.App {
	app := repository.Repository.App.GetByAppId(l.GetDefaultDb(), appid)
	if app == nil {
		return nil
	}

	return app
}

func (l *App) addAppProxy(app *model.App, componentApp *model.App) error {
	return l.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		err := repository.Repository.AppProxy.ClearComponentAppProxy(tx, componentApp)
		if err != nil {
			return err
		}
		appProxy := repository.Repository.AppProxy.AddAppProxy(tx, app, componentApp)
		if appProxy == nil {
			return exception.NewResponseError("设备绑定失败")
		}

		return nil
	})
}

func (l *App) GetSign(app *model.App, params map[string]string) string {
	_, ok := params["sign"]
	if ok {
		delete(params, "sign")
	}

	keys := make([]string, len(params))
	numFieldCount := 0
	for k, _ := range params {
		keys[numFieldCount] = k
		numFieldCount++
	}
	sort.Strings(keys)

	numFieldCount = 0
	paramList := make([]string, len(params))
	for _, k := range keys {
		s := fmt.Sprintf("%s=%v", k, params[k])
		paramList[numFieldCount] = s
		numFieldCount++
	}
	str := strings.Join(paramList, "&") + app.AppSecret

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
