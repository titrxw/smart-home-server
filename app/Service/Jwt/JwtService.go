package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golobby/container/v3/pkg/container"
	base "github.com/titrxw/smart-home-server/app/Service/Base"
)

const JWT_SERVICE = "service:jwt"

type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

type JwtService struct {
	base.ServiceAbstract
	Iss             string
	Subject         string
	Audience        string
	NotBeforeSecond int64
	TTL             int64
	PrivateKey      string
	PublicKey       string
}

func NewJwtService(Iss string, Subject string, Audience string, NotBeforeSecond int64, TTL int64, PrivateKey string, PublicKey string) *JwtService {
	return &JwtService{
		Iss:             Iss,
		Subject:         Subject,
		Audience:        Audience,
		NotBeforeSecond: NotBeforeSecond,
		TTL:             TTL,
		PrivateKey:      PrivateKey,
		PublicKey:       PublicKey,
	}
}

func (jwtService *JwtService) MakeToken(Payload interface{}) (string, error) {
	var aud []string
	if jwtService.Audience != "" {
		aud = append(aud, jwtService.Audience)
	}
	clams := Claims{
		Payload: Payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   jwtService.Iss,
			Subject:  jwtService.Subject,
			Audience: aud,
			//签名生效时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtService.NotBeforeSecond) * time.Second)),
			//发放时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtService.TTL) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, clams)
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(jwtService.PrivateKey))
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func (jwtService *JwtService) ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(jwtService.PublicKey))
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (jwtService *JwtService) ValidateToken(token *jwt.Token) error {
	clams := token.Claims.(Claims)
	if clams.Issuer != jwtService.Iss {
		return errors.New("issued error")
	}
	if clams.Subject != jwtService.Subject {
		return errors.New("subject error")
	}
	if !clams.VerifyAudience(jwtService.Audience, false) {
		return errors.New("audience error")
	}
	if err := clams.Valid(); err != nil {
		return err
	}

	return nil
}

func GetJwtService(container container.Container) *JwtService {
	var service *JwtService
	err := container.NamedResolve(&service, JWT_SERVICE)
	if err != nil {
		panic(err)
	}

	return service
}
