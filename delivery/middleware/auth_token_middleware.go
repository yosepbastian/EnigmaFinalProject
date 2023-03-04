package middleware

import (
	"fmt"
	"kel1-stockbite-projects/utils/authenticator"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("DWici392-sl93wcFD@")
var ApplicationName = "stockbite"

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleWare interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleWare struct {
	acctToken authenticator.AccessToken
}

func NewTokenValidator(acctToken authenticator.AccessToken) AuthTokenMiddleWare {
	return &authTokenMiddleWare{
		acctToken: acctToken,
	}
}

func (a *authTokenMiddleWare) RequireToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login/order" {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized1",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.JSON(401, gin.H{
					"message": "Unauthorized2",
				})
				c.Abort()
				return
			}
			token, err := authenticator.VerifyAccessToken(tokenString)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized3",
				})
				c.Abort()
				return
			}
			fmt.Println(token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Unauthorizedr",
				})
				c.Abort()
				return
			}
		}
	}
}
