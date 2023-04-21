package common

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	"github.com/titrxw/smart-home-server/app/common/emqx"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
	errorhandler "github.com/we7coreteam/w7-rangine-go/src/core/err_handler"
)

type Provider struct {
	provider.Abstract
}

func (p Provider) Register() {
	p.registerEmqxService()
}

func (p Provider) registerEmqxService() {
	client := kernel.NewClient(p.GetConfig().GetString("emqx.host"), p.GetConfig().GetString("emqx.app_id"), p.GetConfig().GetString("emqx.app_secret"))
	messageService := emqx.NewEmqxMessageService(client)
	clientService := emqx.NewEmqxClientService(client)
	err := p.GetContainer().NamedSingleton(emqx.EmqMessageService, func() *emqx.Message {
		return messageService
	})
	if errorhandler.Found(err) {
		panic(err)
	}

	err = p.GetContainer().NamedSingleton(emqx.EmqClientService, func() *emqx.Client {
		return clientService
	})
	if errorhandler.Found(err) {
		panic(err)
	}
}
