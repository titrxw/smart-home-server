package emqx

import (
	"context"

	"github.com/golobby/container/v3/pkg/container"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	acl_entity "github.com/titrxw/emqx-sdk/src/Acl/Entity"
	acl_handler "github.com/titrxw/emqx-sdk/src/Acl/Handler"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	auth_entity "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	auth_handler "github.com/titrxw/emqx-sdk/src/Auth/Handler"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

const EMQ_CLIENT_SERVICE = "service:emq:client"

type EmqxClientService struct {
	EmqxServiceAbstract
	auth *auth.Auth
	acl  *acl.Acl
}

func NewEmqxClientService(EmqxClient *kernel.EmqxClient) *EmqxClientService {
	return &EmqxClientService{
		EmqxServiceAbstract: EmqxServiceAbstract{
			EmqxClient: EmqxClient,
		},
	}
}

func (emqxClientService *EmqxClientService) getAuth() *auth.Auth {
	if emqxClientService.auth == nil {
		authHandler := auth_handler.NewMnesiaAuthHandler(emqxClientService.getEmqxClient())
		emqxClientService.auth = auth.NewAuth(authHandler, nil)
	}

	return emqxClientService.auth
}

func (emqxClientService *EmqxClientService) getAcl() *acl.Acl {
	if emqxClientService.acl == nil {
		aclHandler := acl_handler.NewMnesiaAclHandler(emqxClientService.getEmqxClient())
		emqxClientService.acl = acl.NewAcl(aclHandler)
	}

	return emqxClientService.acl
}

func (emqxClientService *EmqxClientService) AddClient(ctx context.Context, clientId string, password string, salt string) error {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)
	authEntity.SetPassword(password)
	authEntity.SetSalt(salt)

	return emqxClientService.getAuth().Set(ctx, authEntity, false)
}

func (emqxClientService *EmqxClientService) DeleteClient(ctx context.Context, clientId string) error {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)

	return emqxClientService.getAuth().Delete(ctx, authEntity, false)
}

func (emqxClientService *EmqxClientService) AddClientPubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPub()
	aclEntity.SetAccessAllow()

	return emqxClientService.getAcl().Set(ctx, aclEntity, false)
}

func (emqxClientService *EmqxClientService) AddClientSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionSub()
	aclEntity.SetAccessAllow()

	return emqxClientService.getAcl().Set(ctx, aclEntity, false)
}

func (emqxClientService *EmqxClientService) AddClientPubSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPubSub()
	aclEntity.SetAccessAllow()

	return emqxClientService.getAcl().Set(ctx, aclEntity, false)
}

func (emqxClientService *EmqxClientService) DeleteClientAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)

	return emqxClientService.getAcl().Delete(ctx, aclEntity, false)
}

func GetEmqxClientService(container container.Container) *EmqxClientService {
	var service *EmqxClientService
	err := container.NamedResolve(&service, EMQ_CLIENT_SERVICE)
	if err != nil {
		panic(err)
	}

	return service
}
