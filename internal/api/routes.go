package api

import (
	"portfoliosite_v4_admin_auth_service/internal/handlers"
	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/internal/middleware"
	"portfoliosite_v4_admin_auth_service/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gormDB *gorm.DB, jwtManager *jwtmanager.JWTManager) *gin.Engine {
	router := gin.Default()

	repo := repository.NewUserRepository(gormDB)

	// Public routes
	router.POST("/login", handlers.LoginHandler(repo, jwtManager)) // Login Handler

	// Protected routes using JWT middleware
	api := router.Group("/")
	api.Use(middleware.AuthMiddleware(jwtManager, repo))
	{
		api.POST("/users", handlers.CreateUserHandler(repo))  // Create a new user
		api.GET("/users/:id", handlers.GetUserByIDHandler(repo))  // Retrieve a user by ID
		api.PUT("/users/:id", handlers.UpdateUserHandler(repo))  // Update a user by ID
		api.DELETE("/users/:id", handlers.DeleteUserHandler(repo))  // Delete a user by ID
		api.GET("/users", handlers.ListAllUsersHandler(repo))  // List all users
		api.GET("/users/role/:role", handlers.ListUsersByRoleHandler(repo))  // List users by role
	}

	// API Test Routes
	router.GET("/ping", handlers.PingHandler)

	return router
}
