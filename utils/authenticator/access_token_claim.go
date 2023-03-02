package authenticator

import "github.com/golang-jwt/jwt"

type JwtClaims struct{
	jwt.StandardClaims
	Username string `json:"Username"`
	Email string `json:"Email"`
}