package emqx

import (
	"context"
	app "github.com/we7coreteam/w7-rangine-go/src"

	acl "github.com/titrxw/emqx-sdk/src/Acl"
	identity "github.com/titrxw/emqx-sdk/src/Acl/Entity"
	handler "github.com/titrxw/emqx-sdk/src/Acl/Handler"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	authenticity "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	authhandler "github.com/titrxw/emqx-sdk/src/Auth/Handler"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

const EmqClientService = "service:emq:client"

type Client struct {
	Abstract
	auth *auth.Auth
	acl  *acl.Acl
}

func NewEmqxClientService(client *kernel.EmqxClient) *Client {
	return &Client{
		Abstract: Abstract{
			EmqxClient: client,
		},
	}
}

func (s *Client) getAuth() *auth.Auth {
	if s.auth == nil {
		authHandler := authhandler.NewMnesiaAuthHandler(s.getEmqxClient())
		s.auth = auth.NewAuth(authHandler, nil)
	}

	return s.auth
}

func (s *Client) getAcl() *acl.Acl {
	if s.acl == nil {
		aclHandler := handler.NewMnesiaAclHandler(s.getEmqxClient())
		s.acl = acl.NewAcl(aclHandler)
	}

	return s.acl
}

func (s *Client) AddClient(ctx context.Context, clientId string, password string, salt string) error {
	authEntity := new(authenticity.AuthEntity)
	authEntity.SetClientName(clientId)
	authEntity.SetPassword(password)
	authEntity.SetSalt(salt)

	return s.getAuth().Set(ctx, authEntity, false)
}

func (s *Client) DeleteClient(ctx context.Context, clientId string) error {
	authEntity := new(authenticity.AuthEntity)
	authEntity.SetClientName(clientId)

	return s.getAuth().Delete(ctx, authEntity, false)
}

func (s *Client) AddClientPubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(identity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPub()
	aclEntity.SetAccessAllow()

	return s.getAcl().Set(ctx, aclEntity, false)
}

func (s *Client) AddClientSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(identity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionSub()
	aclEntity.SetAccessAllow()

	return s.getAcl().Set(ctx, aclEntity, false)
}

func (s *Client) AddClientPubSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(identity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPubSub()
	aclEntity.SetAccessAllow()

	return s.getAcl().Set(ctx, aclEntity, false)
}

func (s *Client) DeleteClientAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(identity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)

	return s.getAcl().Delete(ctx, aclEntity, false)
}

func GetEmqxClientService() *Client {
	var service *Client
	err := app.GApp.GetContainer().NamedResolve(&service, EmqClientService)
	if err != nil {
		panic(err)
	}

	return service
}
