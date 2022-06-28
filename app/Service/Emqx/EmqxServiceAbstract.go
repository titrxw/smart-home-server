package emqx

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	base "github.com/titrxw/smart-home-server/app/Service/Base"
)

type EmqxServiceAbstract struct {
	base.ServiceAbstract
	EmqxClient *kernel.EmqxClient
}

func (this *EmqxServiceAbstract) getEmqxClient() *kernel.EmqxClient {
	return this.EmqxClient
}
