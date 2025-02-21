package jwtmanager

import (
	"fmt"
	"time"

	"portfoliosite_v4_admin_auth_service/internal/repository"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
    secretKey string
    tokenDuration time.Duration
}

type Claims struct {
    jwt.StandardClaims
    Username string `json:"username"`  
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
    return &JWTManager{
        secretKey: secretKey,
        tokenDuration: tokenDuration,
    }
}

// GenerateToken creates a new JWT token for a given user ID
func (manager *JWTManager) GenerateToken(userID string, username string) (string, error) {
    claims := &Claims{
        StandardClaims: jwt.StandardClaims{
            Subject:   userID,
            ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
        },
        Username: username,  // Add the username to the token
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) VerifyToken(tokenStr string, userRepository *repository.UserRepository) (*jwt.Token, *Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(manager.secretKey), nil
    })

    if err != nil {
        return nil, nil, err
    }

    if !token.Valid {
        return nil, nil, fmt.Errorf("invalid token")
    }

    // Retrieve the user from the database
    user, err := userRepository.GetUserByID(claims.Subject)
    if err != nil || user == nil {
        return nil, nil, fmt.Errorf("no such user found")
    }

    // Optionally check if the username matches
    if user.Username != claims.Username {
        return nil, nil, fmt.Errorf("invalid user information")
    }

    return token, claims, nil
}
