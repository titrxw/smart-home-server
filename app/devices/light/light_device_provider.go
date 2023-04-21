package light

import (
	"github.com/titrxw/smart-home-server/app/devices/manager"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
)

type Provider struct {
	provider.Abstract
}

func (p *Provider) Register() {
	manager.RegisterDevice(new(DeviceAdapter))
}
