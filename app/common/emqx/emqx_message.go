package emqx

import (
	"context"
	app "github.com/we7coreteam/w7-rangine-go/src"

	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	openapi "github.com/titrxw/emqx-sdk/src/OpenApi"
)

const EmqMessageService = "service:emq:message"

type Message struct {
	Abstract
	openapiFactory *openapi.OpenApiFactory
}

func NewEmqxMessageService(EmqxClient *kernel.EmqxClient) *Message {
	return &Message{
		Abstract: Abstract{
			EmqxClient: EmqxClient,
		},
	}
}

func (s *Message) getOpenApiFactory() *openapi.OpenApiFactory {
	if s.openapiFactory == nil {
		s.openapiFactory = openapi.NewOpenApiFactory(s.getEmqxClient())
	}

	return s.openapiFactory
}

func (s *Message) Publish(ctx context.Context, clientId string, topic string, payload string, qos int, retain bool) error {
	return s.getOpenApiFactory().Message().Publish(ctx, clientId, topic, payload, qos, retain)
}

func GetEmqxMessageService() *Message {
	var service *Message
	err := app.GApp.GetContainer().NamedResolve(&service, EmqMessageService)
	if err != nil {
		panic(err)
	}

	return service
}
