package handlers

import (
	"net/http"

	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(repo *repository.UserRepository, jwtManager *jwtmanager.JWTManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
            return
        }

        user, err := repo.GetUserByEmail(input.Email)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
            return
        }

        token, err := jwtManager.GenerateToken(user.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
    }
}
