package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type AppLogic struct {
	LogicAbstract
}

func (appLogic *AppLogic) GetAppByAppId(appid string) *model.App {
	app := repository.Repository.AppRepository.GetByAppId(appLogic.GetDefaultDb(), appid)
	if app == nil {
		return nil
	}

	return app
}

func (appLogic *AppLogic) addAppProxy(app *model.App, componentApp *model.App) error {
	return appLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		err := repository.Repository.AppProxyRepository.ClearComponentAppProxy(tx, componentApp)
		if err != nil {
			return err
		}
		appProxy := repository.Repository.AppProxyRepository.AddAppProxy(tx, app, componentApp)
		if appProxy == nil {
			return exception.NewLogicError("设备绑定失败")
		}

		return nil
	})
}

func (appLogic *AppLogic) GetSign(app *model.App, params map[string]string) string {
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
