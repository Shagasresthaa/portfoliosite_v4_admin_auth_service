package api

import (
	"portfoliosite_v4_admin_auth_service/internal/handlers"
	"portfoliosite_v4_admin_auth_service/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gormDB *gorm.DB) *gin.Engine {
    router := gin.Default()

    repo := repository.NewUserRepository(gormDB)

    // User routes
    router.POST("/users", handlers.CreateUserHandler(repo))  // Create a new user
    router.GET("/users/:id", handlers.GetUserByIDHandler(repo))  // Retrieve a user by ID
    router.PUT("/users/:id", handlers.UpdateUserHandler(repo))  // Update a user by ID
    router.DELETE("/users/:id", handlers.DeleteUserHandler(repo))  // Delete a user by ID
    router.GET("/users", handlers.ListAllUsersHandler(repo))  // List all users
    router.GET("/users/role/:role", handlers.ListUsersByRoleHandler(repo))  // List users by role

    // API Test Routes
    router.GET("/ping", handlers.PingHandler)

    return router
}
