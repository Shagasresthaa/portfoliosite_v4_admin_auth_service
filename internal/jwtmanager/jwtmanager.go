package jwtmanager

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
    secretKey     string
    tokenDuration time.Duration
}

type Claims struct {
    jwt.StandardClaims
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
    return &JWTManager{
        secretKey:     secretKey,
        tokenDuration: tokenDuration,
    }
}

func (manager *JWTManager) GenerateToken(userID string) (string, error) {
    claims := &Claims{
        StandardClaims: jwt.StandardClaims{
            Subject:   userID,
            ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) VerifyToken(tokenStr string) (*jwt.Token, error) {
    return jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(manager.secretKey), nil
    })
}
