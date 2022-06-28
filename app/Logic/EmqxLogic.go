package logic

import (
	"context"
	global "github.com/titrxw/go-framework/src/Global"
	"strconv"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	emqx "github.com/titrxw/smart-home-server/app/Service/Emqx"
)

type EmqxLogic struct {
	LogicAbstract
}

func (this EmqxLogic) AddEmqxClient(ctx context.Context, device *model.Device) error {
	emqxService := emqx.GetEmqxClientService(global.FApp.Container)
	err := emqxService.AddClient(ctx, device.App.AppId, device.App.AppSecret, "")
	if err != nil {
		return err
	}

	err = emqxService.AddClientSubAcl(ctx, device.App.AppId, this.GetClientSubTopic(device.App.AppId))
	if err != nil {
		return err
	}
	err = emqxService.AddClientPubAcl(ctx, device.App.AppId, this.GetClientPubTopic(device.App.AppId))
	if err != nil {
		return err
	}

	return nil
}

func (this EmqxLogic) DeleteEmqxClient(ctx context.Context, device *model.Device) error {
	emqxService := emqx.GetEmqxClientService(global.FApp.Container)
	err := emqxService.DeleteClient(ctx, device.App.AppId)
	if err != nil {
		return err
	}

	err = emqxService.DeleteClientAcl(ctx, device.App.AppId, this.GetClientSubTopic(device.App.AppId))
	if err != nil {
		return err
	}
	err = emqxService.DeleteClientAcl(ctx, device.App.AppId, this.GetClientPubTopic(device.App.AppId))
	if err != nil {
		return err
	}

	return nil
}

func (this EmqxLogic) PubClientOperate(ctx context.Context, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	if deviceOperateLog.OperatePayload == nil {
		deviceOperateLog.OperatePayload = make(model.OperatePayload)
	}
	deviceOperateLog.OperatePayload["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	deviceOperateLog.OperatePayload["nonce"] = helper.RandomStr(12)

	message := cloudevents.NewEvent()
	message.SetID(deviceOperateLog.OperateNumber)
	message.SetSource(deviceOperateLog.Source)
	message.SetType(deviceOperateLog.OperateName)
	message.SetSubject("iot_device")
	message.SetTime(time.Time(deviceOperateLog.OperateTime))
	err := message.SetData(cloudevents.ApplicationJSON, deviceOperateLog.OperatePayload)
	if err != nil {
		return err
	}
	tmpByte, err := message.MarshalJSON()
	if err != nil {
		return err
	}

	retain := false
	if deviceOperateLog.OperateLevel > 1 {
		retain = true
	}

	return emqx.GetEmqxMessageService(global.FApp.Container).Publish(ctx, device.App.AppId, EmqxLogic{}.GetClientSubTopic(device.App.AppId), string(tmpByte), int(deviceOperateLog.OperateLevel), retain)
}

func (this EmqxLogic) GetClientSubTopic(appId string) string {
	return "/iot/" + global.FApp.Name + "/device/" + appId + "/ctrl"
}

func (this EmqxLogic) GetClientPubTopic(appId string) string {
	return "/iot/" + global.FApp.Name + "/device/" + appId + "/reply"
}
