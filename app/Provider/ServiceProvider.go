package provider

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	emqx "github.com/titrxw/smart-home-server/app/Service/Emqx"
	jwt "github.com/titrxw/smart-home-server/app/Service/Jwt"
	utils "github.com/titrxw/smart-home-server/app/Utils"
	"github.com/titrxw/smart-home-server/config"
)

type ServiceProvider struct {
	provider.ProviderAbstract
}

func (this *ServiceProvider) Register(options interface{}) {
	config := options.(*config.Config)
	this.registerJwtService(config)
	this.registerEmqxService(config)
}

func (this *ServiceProvider) registerJwtService(options *config.Config) {
	this.RegisterAutoPanic(jwt.JWT_SERVICE, func() *jwt.JwtService {
		var err error
		options.Jwt.PrivateKey, err = utils.Decrypt(options.Jwt.PrivateKey, options.Common.SecureKey)
		if err != nil {
			panic(err)
		}
		options.Jwt.PublicKey, err = utils.Decrypt(options.Jwt.PublicKey, options.Common.SecureKey)
		if err != nil {
			panic(err)
		}

		return jwt.NewJwtService(
			options.Jwt.Iss,
			options.Jwt.Subject,
			options.Jwt.Audience,
			options.Jwt.NotBeforeSecond,
			options.Jwt.TTL,
			options.Jwt.PrivateKey,
			options.Jwt.PublicKey,
		)
	})
}

func (this *ServiceProvider) registerEmqxService(options *config.Config) {
	client := kernel.NewClient(options.Emqx.Host, options.Emqx.AppId, options.Emqx.AppSecret)
	this.RegisterAutoPanic(emqx.EMQ_MESSAGE_SERVICE, func() *emqx.EmqxMessageService {
		return emqx.NewEmqxMessageService(client)
	})
	this.RegisterAutoPanic(emqx.EMQ_CLIENT_SERVICE, func() *emqx.EmqxClientService {
		return emqx.NewEmqxClientService(client)
	})
}
