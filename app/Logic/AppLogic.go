package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AppLogic struct {
	LogicAbstract
}

func (appLogic *AppLogic) GetAppAttachPath(appId string, ext string) (string, error) {
	dir := "./public/upload/img/" + appId + "/" + time.Now().Format("20060102")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", exception.NewLogicError("创建文件夹失败")
	}

	fileUnixName := strconv.FormatInt(time.Now().UnixMicro(), 10) + ext

	return path.Join(dir, fileUnixName), nil
}

func (appLogic *AppLogic) GetRemoteFileUrlByAttachPath(filePath string) string {
	return strings.NewReplacer("public/upload", "").Replace(filePath)
}

func (appLogic *AppLogic) GetAppByAppId(appid string) *model.App {
	app := repository.AppRepository{}.GetByAppId(appLogic.GetDefaultDb(), appid)
	if app == nil {
		return nil
	}

	return app
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
