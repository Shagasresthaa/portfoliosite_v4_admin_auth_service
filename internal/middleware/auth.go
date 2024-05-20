package middleware

import (
	"net/http"
	"strings"

	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager *jwtmanager.JWTManager) gin.HandlerFunc {
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
        token, err := jwtManager.VerifyToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        if claims, ok := token.Claims.(*jwtmanager.Claims); ok && token.Valid {
            // Attach claims to the context for use in the handler
            c.Set("userID", claims.Subject)
            c.Next()
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        }
    }
}
