package middleware

import (
	"kel1-stockbite-projects/utils/authenticator"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	return func(ctx *gin.Context) {
		h := authHeader{}

		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return

		}

		token, err := a.acctToken.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		if token != nil {
			ctx.Next()
		} else {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return

		}

	}

}


