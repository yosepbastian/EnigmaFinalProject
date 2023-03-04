package authenticator

import (
	"fmt"
	"kel1-stockbite-projects/config"
	"kel1-stockbite-projects/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessToken interface {
	CreateAccessToken(cred *models.UserLogin) (string, error)
}
type accessToken struct {
	conf config.TokenConfig
}
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("DWici392-sl93wcFD@")
var ApplicationName = "stockbite"



func (t *accessToken) CreateAccessToken(cred *models.UserLogin) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.conf.AccessTokenLifeTime)

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},

		Email:    cred.Email,
		Password: cred.Password,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)
	return token.SignedString(JwtSignatureKey)
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JwtSignatureKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		conf: config,
	}
}
