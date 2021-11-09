package emqx

import (
	"context"
	openapi "github.com/titrxw/emqx-sdk/src/OpenApi"
)

type EmqxMessageService struct {
	EmqxServiceAbstract
	openapiFactory *openapi.OpenApiFactory
}

func (this *EmqxMessageService) getOpenApiFactory() *openapi.OpenApiFactory {
	if this.openapiFactory == nil {
		this.openapiFactory = openapi.NewOpenApiFactory(this.getEmqxClient())
	}

	return this.openapiFactory
}

func (this *EmqxMessageService) Publish(ctx context.Context, topic string, clientId string, payload string) (bool, error) {
	return this.getOpenApiFactory().Message().Publish(ctx, topic, clientId, payload, 2, true)
}

var GEmqxMessageService = new(EmqxMessageService)
