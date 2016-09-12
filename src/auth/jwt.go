package auth

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTAuthenticationBackend struct {
	signingKey []byte
	expDelta   int
}

var authBackend *JWTAuthenticationBackend = nil

func NewJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackend != nil {
		return authBackend
	}

	authBackend := &JWTAuthenticationBackend{
		signingKey: []byte("Flzx3000c"),
		expDelta:   72,
	}

	return authBackend
}

func (backend *JWTAuthenticationBackend) GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(backend.expDelta)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = userId
	token.Claims = claims
	tokenString, err := token.SignedString(backend.signingKey)
	if err != nil {
		log.Printf("[ERROR] Signing token failed: %s", err)
		return "", err
	}

	return tokenString, nil
}
