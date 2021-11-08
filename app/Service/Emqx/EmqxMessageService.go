package emqx

import (
	"context"
	openapi "github.com/titrxw/emqx-sdk/src/OpenApi"
)

type EmqxMessageService struct {
	EmqxServiceAbstract
}

func (this *EmqxServiceAbstract) getOpenApiFactory() *openapi.OpenApiFactory {
	return openapi.NewOpenApiFactory(this.getEmqxClient())
}

func (this *EmqxMessageService) Publish(ctx context.Context, topic string, clientId string, payload string) (bool, error) {
	return this.getOpenApiFactory().Message().Publish(ctx, topic, clientId, payload, 2, true)
}

var GEmqxMessageService = new(EmqxMessageService)
