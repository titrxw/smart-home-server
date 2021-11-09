package emqx

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	base "github.com/titrxw/smart-home-server/app/Service/Base"
	"github.com/titrxw/smart-home-server/config"
)

type EmqxServiceAbstract struct {
	base.ServiceAbstract
	EmqxConfig config.Emqx
}

func (this *EmqxServiceAbstract) getEmqxClient() *kernel.EmqxClient {
	return kernel.NewClient(this.EmqxConfig.Host, this.EmqxConfig.AppId, this.EmqxConfig.AppSecret)
}
