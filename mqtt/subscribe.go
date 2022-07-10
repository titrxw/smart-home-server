package mqtt

import (
	"github.com/titrxw/smart-home-server/app/Mqtt/Subscribe"
	"strconv"

	"github.com/titrxw/smart-home-server/app"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	mqtt "github.com/titrxw/smart-home-server/app/Mqtt"
)

func RegisterSubscribe(app *app.App) {
	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReplaySubscribe(logic.Logic.EmqxLogic.GetClientOperatePubTopic("+")))
	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe(logic.Logic.EmqxLogic.GetClientReportTopic("+")))
	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceStatusChangeSubscribe("$SYS/brokers/+/clients/+/+"))

	go mqtt.GetSubscribeManager().Start(app.Config.Mqtt.Host, strconv.Itoa(app.Config.Mqtt.Port), app.Config.Mqtt.UserName, app.Config.Mqtt.Password)
}
