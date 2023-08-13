package logic

import (
	"context"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/emqx"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
)

type Emqx struct {
	logic.Abstract

	emqxHttpClient    *emqx.Client
	emqxMessageClient *emqx.Message
}

func (l *Emqx) getEmqxHttpClient() *emqx.Client {
	if l.emqxHttpClient == nil {
		l.emqxHttpClient = emqx.NewEmqxClientService(
			kernel.NewClient(facade.GetConfig().GetString("emqx.host"), facade.GetConfig().GetString("emqx.app_id"), facade.GetConfig().GetString("emqx.app_secret")),
		)
	}

	return l.emqxHttpClient
}

func (l *Emqx) getEmqxMessageClient() *emqx.Message {
	if l.emqxMessageClient == nil {
		l.emqxMessageClient = emqx.NewEmqxMessageService(
			kernel.NewClient(facade.GetConfig().GetString("emqx.host"), facade.GetConfig().GetString("emqx.app_id"), facade.GetConfig().GetString("emqx.app_secret")),
		)
	}

	return l.emqxMessageClient
}

func (l *Emqx) AddEmqxClient(ctx context.Context, device *model.Device, topicAcl map[string][]string) error {
	err := l.getEmqxHttpClient().AddClient(ctx, device.App.AppId, device.App.AppSecret, "")
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
					err = l.getEmqxHttpClient().AddClientPubAcl(ctx, device.App.AppId, topic)
					if err != nil {
						return err
					}
				}
				if acl == "sub" {
					err = l.getEmqxHttpClient().AddClientSubAcl(ctx, device.App.AppId, topic)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (l *Emqx) DeleteEmqxClient(ctx context.Context, device *model.Device) error {
	return l.getEmqxHttpClient().DeleteClient(ctx, device.App.AppId)
}

func (l *Emqx) PubClientOperate(ctx context.Context, appId string, topic string, payload string, level int) error {
	retain := false
	if level > 1 {
		retain = true
	}
	return l.getEmqxMessageClient().Publish(ctx, appId, topic, payload, level, retain)
}
