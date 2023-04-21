package jwt

import (
	"errors"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const ServiceName = "service:jwt"

type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

type Service struct {
	Iss             string
	Subject         string
	Audience        string
	NotBeforeSecond int64
	TTL             int64
	PrivateKey      string
	PublicKey       string
}

func NewJwtService(Iss string, Subject string, Audience string, NotBeforeSecond int64, TTL int64, PrivateKey string, PublicKey string) *Service {
	return &Service{
		Iss:             Iss,
		Subject:         Subject,
		Audience:        Audience,
		NotBeforeSecond: NotBeforeSecond,
		TTL:             TTL,
		PrivateKey:      PrivateKey,
		PublicKey:       PublicKey,
	}
}

func (s *Service) MakeToken(Payload interface{}) (string, error) {
	var aud []string
	if s.Audience != "" {
		aud = append(aud, s.Audience)
	}
	clams := Claims{
		Payload: Payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   s.Iss,
			Subject:  s.Subject,
			Audience: aud,
			//签名生效时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(s.NotBeforeSecond) * time.Second)),
			//发放时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.TTL) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, clams)
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(s.PrivateKey))
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func (s *Service) ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(s.PublicKey))
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) ValidateToken(token *jwt.Token) error {
	clams := token.Claims.(Claims)
	if clams.Issuer != s.Iss {
		return errors.New("issued error")
	}
	if clams.Subject != s.Subject {
		return errors.New("subject error")
	}
	if !clams.VerifyAudience(s.Audience, false) {
		return errors.New("audience error")
	}
	if err := clams.Valid(); err != nil {
		return err
	}

	return nil
}

func GetJwtService() *Service {
	var service *Service
	err := app.GApp.GetContainer().NamedResolve(&service, ServiceName)
	if err != nil {
		panic(err)
	}

	return service
}
