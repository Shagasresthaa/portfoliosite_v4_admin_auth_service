package api

import (
	"portfoliosite_v4_admin_auth_service/internal/handlers"
	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/internal/middleware"
	"portfoliosite_v4_admin_auth_service/internal/repository"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gormDB *gorm.DB, jwtManager *jwtmanager.JWTManager) *gin.Engine {
	router := gin.Default()

	// Configure CORS settings.
    router.Use(cors.New(cors.Config{
        // Only allow these origins to access the API.
        AllowOrigins: []string{
            "https://www.sresthaa.com",
            "http://localhost:3000",
            "http://localhost:3001",
            "http://localhost:4000",
        },
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	repo := repository.NewUserRepository(gormDB)

	// Public routes
	router.POST("/api/admin/login", handlers.LoginHandler(repo, jwtManager)) // Login Handler
	// Note to self: you will find this useful when you deploy the first time or you lock yourself out and need to recreate your user maybe because you nuked the db
	// For the love of god dont you make this public and if you did donot forget to remove it and add it back to jwtmanager
	// router.POST("/users", handlers.CreateUserHandler(repo)) 

	// Protected routes using JWT middleware
	api := router.Group("/api/admin")
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
	router.GET("/api/admin/ping", handlers.PingHandler)

	return router
}
