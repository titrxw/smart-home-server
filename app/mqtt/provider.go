package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	"github.com/titrxw/smart-home-server/app/mqtt/subscribe"
	"github.com/titrxw/smart-home-server/app/pkg/emqx"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type Provider struct {
	server *Server
}

func (p *Provider) Register(config *viper.Viper, serverFactory server.Manager) *Provider {
	server := NewMqttSubServer(config, &SubscribeManager{
		subscribeMap: make(map[string]mqtt.MessageHandler),
		emqxHttpClient: emqx.NewEmqxClientService(
			kernel.NewClient(config.GetString("emqx.host"), config.GetString("emqx.app_id"), config.GetString("emqx.app_secret")),
		),
	})
	serverFactory.RegisterServer(server)

	p.server = server
	p.RegisterMqttSubscribe()

	return p
}

func (p *Provider) RegisterMqttSubscribe() {
	p.server.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe("/iot/" + helper.GetAppName() + "/device/+/report"))
	p.server.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceStatusChangeSubscribe("$SYS/brokers/+/clients/+/+"))
}

func (p *Provider) Export() *Server {
	return p.server
}
