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
        var loginDetails struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.BindJSON(&loginDetails); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
            return
        }

        user, err := repo.GetUserByEmail(loginDetails.Email)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed, user not found"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed, incorrect password"})
            return
        }

        accessToken, err := jwtManager.GenerateToken(user.ID, user.Username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "access_token": accessToken,
            "role": user.Role,  
        })
    }
}
