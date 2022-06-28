package tests

import (
	"github.com/titrxw/smart-home-server/app"
	service "github.com/titrxw/smart-home-server/app/Service/Jwt"
	"testing"
)

func TestJwt(t *testing.T) {
	t.Run("testJwt", func(t *testing.T) {
		payload := map[string]string{
			"test": "test",
		}

		app.GApp.Config.Jwt.PrivateKey = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIAh5qA3rmqQQuu0vbKV/+zouz/y/Iy2pLpIcWUSyImSwoAoGCCqGSM49\nAwEHoUQDQgAEYD54V/vp+54P9DXarYqx4MPcm+HKRIQzNasYSoRQHQ/6S6Ps8tpM\ncT+KvIIC8W/e9k0W7Cm72M1P9jU7SLf/vg==\n-----END EC PRIVATE KEY-----"
		var GJwtService = &service.JwtService{
			Iss:             app.GApp.Config.Jwt.Iss,
			Subject:         app.GApp.Config.Jwt.Subject,
			Audience:        app.GApp.Config.Jwt.Audience,
			NotBeforeSecond: app.GApp.Config.Jwt.NotBeforeSecond,
			TTL:             app.GApp.Config.Jwt.TTL,
			PrivateKey:      app.GApp.Config.Jwt.PrivateKey,
			PublicKey:       app.GApp.Config.Jwt.PublicKey,
		}
		token, _ := GJwtService.MakeToken(payload)

		app.GApp.Config.Jwt.PublicKey = "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYD54V/vp+54P9DXarYqx4MPcm+HK\nRIQzNasYSoRQHQ/6S6Ps8tpMcT+KvIIC8W/e9k0W7Cm72M1P9jU7SLf/vg==\n-----END PUBLIC KEY-----"
		var GJwtService1 = &service.JwtService{
			Iss:             app.GApp.Config.Jwt.Iss,
			Subject:         app.GApp.Config.Jwt.Subject,
			Audience:        app.GApp.Config.Jwt.Audience,
			NotBeforeSecond: app.GApp.Config.Jwt.NotBeforeSecond,
			TTL:             app.GApp.Config.Jwt.TTL,
			PrivateKey:      app.GApp.Config.Jwt.PrivateKey,
			PublicKey:       app.GApp.Config.Jwt.PublicKey,
		}
		payload1, _ := GJwtService1.ParseToken(token)
		payload2 := payload1.Claims.(service.Claims).Payload.(map[string]string)

		_, err := payload2["Test"]
		if err {
			t.Failed()
		}
		t.Skipped()
	})
}
