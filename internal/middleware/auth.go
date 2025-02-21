package middleware

import (
	"net/http"
	"strings"

	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/internal/repository" // Ensure this import path is correct

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager *jwtmanager.JWTManager, userRepository *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            return
        }

        // Expected header: "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
            return
        }

        tokenStr := parts[1]
        token, claims, err := jwtManager.VerifyToken(tokenStr, userRepository)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
            return
        }

        if token.Valid {
            c.Set("userID", claims.Subject)
            c.Next()
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        }
    }
}
