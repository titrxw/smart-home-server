package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/common/emqx"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
)

type Emqx struct {
	Abstract
}

func (l Emqx) AddEmqxClient(ctx context.Context, device *model.Device, topicAcl map[string][]string) error {
	emqxService := emqx.GetEmqxClientService()
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

func (l Emqx) DeleteEmqxClient(ctx context.Context, device *model.Device) error {
	emqxService := emqx.GetEmqxClientService()
	return emqxService.DeleteClient(ctx, device.App.AppId)
}

func (l Emqx) PubClientOperate(ctx context.Context, appId string, topic string, payload string, level int) error {
	retain := false
	if level > 1 {
		retain = true
	}
	return emqx.GetEmqxMessageService().Publish(ctx, appId, topic, payload, level, retain)
}
