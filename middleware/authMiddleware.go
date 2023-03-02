package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireAuth(ctx *gin.Context) {
	fmt.Println("in Midleware")

	ctx.Next()
}
