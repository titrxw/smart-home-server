package mqtt

import (
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type Provider struct {
	provider.Abstract
}

func (p Provider) Register() {
	server.RegisterServer(NewMqttSubServer(p.GetConfig()))
}
