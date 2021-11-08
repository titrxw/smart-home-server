package emqx

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	app "github.com/titrxw/smart-home-server"
	base "github.com/titrxw/smart-home-server/app/Service/Base"
)

type EmqxServiceAbstract struct {
	base.ServiceAbstract
}

func (this *EmqxServiceAbstract) getEmqxClient() *kernel.EmqxClient {
	return kernel.NewClient(app.GApp.Config.Emqx.Host, app.GApp.Config.Emqx.AppId, app.GApp.Config.Emqx.AppSecret)
}
