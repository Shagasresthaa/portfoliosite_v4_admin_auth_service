package jwtmanager

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
    secretKey string
    tokenDuration time.Duration
}

type Claims struct {
    jwt.StandardClaims
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
    return &JWTManager{
        secretKey: secretKey,
        tokenDuration: tokenDuration,
    }
}

// GenerateToken creates a new JWT token for a given user ID
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

// VerifyToken checks the validity of the token
func (manager *JWTManager) VerifyToken(tokenStr string) (*jwt.Token, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(manager.secretKey), nil
    })
    return token, err
}
