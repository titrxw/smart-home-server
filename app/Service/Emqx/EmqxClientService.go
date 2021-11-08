package emqx

import (
	"context"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	acl_entity "github.com/titrxw/emqx-sdk/src/Acl/Entity"
	acl_handler "github.com/titrxw/emqx-sdk/src/Acl/Handler"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	encrypt "github.com/titrxw/emqx-sdk/src/Auth/Encrypt"
	auth_entity "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	auth_handler "github.com/titrxw/emqx-sdk/src/Auth/Handler"
)

type EmqxClientService struct {
	EmqxServiceAbstract
}

func (this *EmqxClientService) getAuthHandler() *auth.Auth {
	authHandler := auth_handler.NewMnesiaAuthHandler(this.getEmqxClient())
	return auth.NewAuth(authHandler, new(encrypt.Sha256SaltEncrypt))
}

func (this *EmqxClientService) getAclHandler() *acl.Acl {
	aclHandler := acl_handler.NewMnesiaAclHandler(this.getEmqxClient())
	return acl.NewAcl(aclHandler)
}

func (this *EmqxClientService) AddClient(ctx context.Context, clientId string, password string, salt string) (bool, error) {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)
	authEntity.SetPassword(password)
	authEntity.SetSalt(salt)

	return this.getAuthHandler().Set(ctx, authEntity, true)
}

func (this *EmqxClientService) AddClientPubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPub()
	aclEntity.SetAccessAllow()

	return this.getAclHandler().Set(ctx, aclEntity, true)
}

func (this *EmqxClientService) AddClientSubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionSub()
	aclEntity.SetAccessAllow()

	return this.getAclHandler().Set(ctx, aclEntity, true)
}

func (this *EmqxClientService) AddClientPubSubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPubSub()
	aclEntity.SetAccessAllow()

	return this.getAclHandler().Set(ctx, aclEntity, true)
}

var GEmqxClientService = new(EmqxClientService)
