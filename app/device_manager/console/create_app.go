package console

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/we7coreteam/w7-rangine-go/src/console"
)

type CreateAppCommand struct {
	console.Abstract
}

func (c CreateAppCommand) GetName() string {
	return "create:app"
}

func (c CreateAppCommand) GetDescription() string {
	return "create app"
}

func (c CreateAppCommand) Configure(command *cobra.Command) {
	command.Flags().Uint8("app-type", 1, "app type. 1:device app,2: open app")
	command.MarkFlagRequired("app-type")
}

func (c CreateAppCommand) Handle(cmd *cobra.Command, args []string) {
	appType, err := cmd.Flags().GetUint8("app-type")
	if err != nil {
		color.Errorln(err)
		return
	}

	app := logic.Logic.App.CreateApp(appType)
	if app == nil {
		color.Errorln("create app fail")
		return
	}

	color.Successln("create app success!")
	color.Red.Printf(" appid: %s, appsecret: %s", app.AppId, app.AppSecret)
}
