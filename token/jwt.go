package token

import (
	"fmt"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var ApplicationName = "ENIGMA"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("P@ssw0rd")

type MyClaims struct {
	jwt.StandardClaims
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"userPassword"`
}

// func main() {
// 	r := gin.Default()
// 	r.Use(AuthTokenMiddleware())
// 	publicRoute := r.Group("/enigma")
// 	publicRoute.POST("/auth", func(c *gin.Context) {
// 		var user Credential
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": "can't bind struct",
// 			})
// 			return
// 		})

// 	publicRoute.GET("/user", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "user",
// 		})
// 	})
// 	err := r.Run("localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}

// }

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/enigma/auth" {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			token, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			fmt.Println(token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
		}
	}
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JwtSignatureKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
func GenerateToken(email string, password string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},
		Email:    email,
		Password: password,
	}
	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)
	return token.SignedString(JwtSignatureKey)
}
