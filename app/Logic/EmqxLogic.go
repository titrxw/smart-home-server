package logic

import (
	"context"
	global "github.com/titrxw/go-framework/src/Global"

	model "github.com/titrxw/smart-home-server/app/Model"
	emqx "github.com/titrxw/smart-home-server/app/Service/Emqx"
)

type EmqxLogic struct {
	LogicAbstract
}

func (emqxLogic EmqxLogic) AddEmqxClient(ctx context.Context, device *model.Device, topicAcl map[string][]string) error {
	emqxService := emqx.GetEmqxClientService(global.FApp.Container)
	err := emqxService.AddClient(ctx, device.App.AppId, device.App.AppSecret, "")
	if err != nil {
		return err
	}

	if topicAcl != nil {
		for acl, topics := range topicAcl {
			for _, topic := range topics {
				if topic == "" {
					continue
				}

				if acl == "pub" {
					err = emqxService.AddClientPubAcl(ctx, device.App.AppId, topic)
					if err != nil {
						return err
					}
				}
				if acl == "sub" {
					err = emqxService.AddClientSubAcl(ctx, device.App.AppId, topic)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (emqxLogic EmqxLogic) DeleteEmqxClient(ctx context.Context, device *model.Device) error {
	emqxService := emqx.GetEmqxClientService(global.FApp.Container)
	return emqxService.DeleteClient(ctx, device.App.AppId)
}

func (emqxLogic EmqxLogic) PubClientOperate(ctx context.Context, appId string, topic string, payload string, level int) error {
	retain := false
	if level > 1 {
		retain = true
	}
	return emqx.GetEmqxMessageService(global.FApp.Container).Publish(ctx, appId, topic, payload, level, retain)
}
