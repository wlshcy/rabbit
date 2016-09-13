package auth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Get("X-AUTH-TOKEN")) == 0 {
			ctx.AbortWithStatus(401)
		}
		tokenString := ctx.Request.Header.Get("X-AUTH-TOKEN")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return "Flzx3000c", nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatus(401)
		}

		ctx.Next()
	}
}
