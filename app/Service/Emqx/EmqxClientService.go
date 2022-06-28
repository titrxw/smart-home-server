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

func (this *EmqxClientService) getAuth() *auth.Auth {
	if this.auth == nil {
		authHandler := auth_handler.NewMnesiaAuthHandler(this.getEmqxClient())
		this.auth = auth.NewAuth(authHandler, nil)
	}

	return this.auth
}

func (this *EmqxClientService) getAcl() *acl.Acl {
	if this.acl == nil {
		aclHandler := acl_handler.NewMnesiaAclHandler(this.getEmqxClient())
		this.acl = acl.NewAcl(aclHandler)
	}

	return this.acl
}

func (this *EmqxClientService) AddClient(ctx context.Context, clientId string, password string, salt string) error {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)
	authEntity.SetPassword(password)
	authEntity.SetSalt(salt)

	return this.getAuth().Set(ctx, authEntity, false)
}

func (this *EmqxClientService) DeleteClient(ctx context.Context, clientId string) error {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)

	return this.getAuth().Delete(ctx, authEntity, false)
}

func (this *EmqxClientService) AddClientPubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, false)
}

func (this *EmqxClientService) AddClientSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionSub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, false)
}

func (this *EmqxClientService) AddClientPubSubAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPubSub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, false)
}

func (this *EmqxClientService) DeleteClientAcl(ctx context.Context, clientId string, topic string) error {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)

	return this.getAcl().Delete(ctx, aclEntity, false)
}

func GetEmqxClientService(container container.Container) *EmqxClientService {
	var service *EmqxClientService
	err := container.NamedResolve(&service, EMQ_CLIENT_SERVICE)
	if err != nil {
		panic(err)
	}

	return service
}
