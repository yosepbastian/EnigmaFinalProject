package authenticator

import (
	"fmt"
	"kel1-stockbite-projects/config"
	"kel1-stockbite-projects/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessToken interface {
	CreateAccessToken(cred *models.UserLogin, name string) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type accessToken struct {
	conf config.TokenConfig
}

func (t *accessToken) CreateAccessToken(cred *models.UserLogin, name string) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.conf.AccessTokenLifeTime)
	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.conf.ApplicationName,
		},

		Name:  name,
		Email: cred.Email,
	}

	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		t.conf.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(t.conf.JwtSignatureKey))
}

func (t accessToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != t.conf.JwtSigningMethod { 
			
			return nil, fmt.Errorf("signing method invalid")
		}

		return t.conf.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims,nil

}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		conf: config,
	}
}
