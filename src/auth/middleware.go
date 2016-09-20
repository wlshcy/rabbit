package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	backend := NewJWTAuthenticationBackend()
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Get("X-AUTH-TOKEN")) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := ctx.Request.Header.Get("X-AUTH-TOKEN")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return backend.signingKey, nil
		})

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			ctx.Set("uid", claims["sub"])
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
