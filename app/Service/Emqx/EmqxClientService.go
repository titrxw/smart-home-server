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
	app "github.com/titrxw/smart-home-server"
)

type EmqxClientService struct {
	EmqxServiceAbstract
	auth *auth.Auth
	acl  *acl.Acl
}

func (this *EmqxClientService) getAuth() *auth.Auth {
	if this.auth == nil {
		authHandler := auth_handler.NewMnesiaAuthHandler(this.getEmqxClient())
		this.auth = auth.NewAuth(authHandler, new(encrypt.Sha256SaltEncrypt))
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

func (this *EmqxClientService) AddClient(ctx context.Context, clientId string, password string, salt string) (bool, error) {
	authEntity := new(auth_entity.AuthEntity)
	authEntity.SetClientName(clientId)
	authEntity.SetPassword(password)
	authEntity.SetSalt(salt)

	return this.getAuth().Set(ctx, authEntity, true)
}

func (this *EmqxClientService) AddClientPubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, true)
}

func (this *EmqxClientService) AddClientSubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionSub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, true)
}

func (this *EmqxClientService) AddClientPubSubAcl(ctx context.Context, clientId string, topic string) (bool, error) {
	aclEntity := new(acl_entity.AclEntity)
	aclEntity.SetClientName(clientId)
	aclEntity.SetTopic(topic)
	aclEntity.SetActionPubSub()
	aclEntity.SetAccessAllow()

	return this.getAcl().Set(ctx, aclEntity, true)
}

var GEmqxClientService = &EmqxClientService{
	EmqxServiceAbstract: EmqxServiceAbstract{
		EmqxConfig: app.GApp.Config.Emqx,
	},
}
