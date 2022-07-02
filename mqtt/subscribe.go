package mqtt

import (
	"strconv"

	"github.com/titrxw/smart-home-server/app"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	mqtt "github.com/titrxw/smart-home-server/app/Mqtt"
)

func RegisterSubscribe(app *app.App) {
	mqtt.GetSubscribeManager().RegisterSubscribe(logic.Logic.EmqxLogic.GetClientOperatePubTopic("+"), mqtt.DeviceReplaySubscribe{}.DeviceReplySubscribe)
	mqtt.GetSubscribeManager().RegisterSubscribe("$SYS/brokers/+/clients/+/+", mqtt.DeviceStatusChangeSubscribe{}.DeviceStatusChangeSubscribe)
	reportTopic := logic.Logic.EmqxLogic.GetClientReportTopic("+")
	mqtt.GetSubscribeManager().RegisterSubscribe(reportTopic, mqtt.NewDeviceReportSubscribe(reportTopic).DeviceReportSubscribe)

	go mqtt.GetSubscribeManager().Start(app.Config.Mqtt.Host, strconv.Itoa(app.Config.Mqtt.Port), app.Config.Mqtt.UserName, app.Config.Mqtt.Password)
}
