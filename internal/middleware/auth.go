package middleware

import (
	"net/http"

	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager *jwtmanager.JWTManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        token, err := jwtManager.VerifyToken(tokenStr)
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }
        c.Next()
    }
}
