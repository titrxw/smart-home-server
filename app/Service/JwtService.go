package service

import (
	"github.com/golang-jwt/jwt/v4"
	base "github.com/titrxw/smart-home-server/app/Service/Base"
	"github.com/titrxw/smart-home-server/config"
	"time"
)

type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

type JwtService struct {
	base.ServiceAbstract
	JwtConfig config.Jwt
}

func (this *JwtService) MakeToken(Payload interface{}) (string, error) {
	clams := Claims{
		Payload: Payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  this.JwtConfig.Iss,
			Subject: this.JwtConfig.Subject,
			//签名生效时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Duration(this.JwtConfig.NotBeforeSecond) * time.Second)),
			//发放时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(this.JwtConfig.TTL) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, clams)
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(this.JwtConfig.PrivateKey))
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

func (this *JwtService) ParseToken(tokenStr string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(this.JwtConfig.PublicKey))
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*Claims).Payload, nil
}
