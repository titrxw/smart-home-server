package emqx

import (
	"context"
	"github.com/golobby/container/v3/pkg/container"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	openapi "github.com/titrxw/emqx-sdk/src/OpenApi"
)

const EMQ_MESSAGE_SERVICE = "service:emq:message"

type EmqxMessageService struct {
	EmqxServiceAbstract
	openapiFactory *openapi.OpenApiFactory
}

func NewEmqxMessageService(EmqxClient *kernel.EmqxClient) *EmqxMessageService {
	return &EmqxMessageService{
		EmqxServiceAbstract: EmqxServiceAbstract{
			EmqxClient: EmqxClient,
		},
	}
}

func (this *EmqxMessageService) getOpenApiFactory() *openapi.OpenApiFactory {
	if this.openapiFactory == nil {
		this.openapiFactory = openapi.NewOpenApiFactory(this.getEmqxClient())
	}

	return this.openapiFactory
}

func (this *EmqxMessageService) Publish(ctx context.Context, clientId string, topic string, payload string, qos int, retain bool) error {
	return this.getOpenApiFactory().Message().Publish(ctx, clientId, topic, payload, qos, retain)
}

func GetEmqxMessageService(container container.Container) *EmqxMessageService {
	var service *EmqxMessageService
	err := container.NamedResolve(&service, EMQ_MESSAGE_SERVICE)
	if err != nil {
		panic(err)
	}

	return service
}
