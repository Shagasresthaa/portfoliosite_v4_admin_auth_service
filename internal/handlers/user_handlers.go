package handlers

import (
	"net/http"
	"portfoliosite_v4_admin_auth_service/internal/models"
	"portfoliosite_v4_admin_auth_service/internal/repository"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ListAllUsersHandler returns all users
func ListAllUsersHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := repo.ListAllUsers()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"users": users})
    }
}

// ListUsersByRoleHandler returns users filtered by role
func ListUsersByRoleHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.Param("role")
        users, err := repo.ListUsersByRole(role)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"users": users})
    }
}

// GetUserByIDHandler returns a user by ID
func GetUserByIDHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        user, err := repo.GetUserByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}

// CreateUserHandler creates a new user
func CreateUserHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input struct {
            Username string
            Email    string
            Password string
            Role     string
			Sponsor  string
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }

        if !validateEmail(input.Email) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }

        newUser := models.NewUser(input.Username, input.Email, string(hashedPassword), input.Role, input.Sponsor)
        if err := repo.CreateUser(newUser); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
            return
        }

        newUser.Password = "" // Do not return the password hash
        c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": newUser})
    }
}

// UpdateUserHandler updates an existing user
func UpdateUserHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var user models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }
        user.ID = id
        if err := repo.UpdateUser(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User updated"})
    }
}

// DeleteUserHandler deletes a user by ID
func DeleteUserHandler(repo *repository.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        if err := repo.DeleteUser(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
    }
}

// Validate email format
func validateEmail(email string) bool {
    regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return regex.MatchString(email)
}
