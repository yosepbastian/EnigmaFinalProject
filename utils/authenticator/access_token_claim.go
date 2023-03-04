package authenticator

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Password string `json:"password"`
}
