package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"github.com/spf13/afero"
	"github.com/titrxw/smart-home-server/router"
	"path"
)

func InitConfig() {
	fs := afero.NewOsFs()
	configFiles, err := afero.Glob(fs, "config/*.yaml")
	if err != nil {
		panic("配置文件缺失")
	}

	config.WithOptions(config.ParseEnv)
	for _, filePath := range configFiles {
		data := make(map[string]interface{})
		content, err := afero.ReadFile(fs, filePath)
		if err != nil {
			panic(err.Error() + " " + filePath + "读取失败")
		}
		err = yaml.Decoder(content, data)
		if err != nil {
			panic(err.Error() + " " + filePath + "解析失败")
		}

		fileName := path.Base(filePath)
		err = config.LoadData(map[string]interface{}{
			fileName[0 : len(fileName)-len(path.Ext(fileName))]: data,
		})
		if err != nil {
			panic(err.Error() + " " + filePath + "配置保存失败")
		}
	}
}

func InitRouter(server *gin.Engine) {
	router.Register(server)
}

func main() {
	InitConfig()
	gin.SetMode(config.DefString("app.app_model", "release"))

	server := gin.Default()
	InitRouter(server)

	server.Run(config.DefString("app.app_host") + ":" + config.DefString("app.app_port"))
}
