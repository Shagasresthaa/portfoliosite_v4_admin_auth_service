package main

import (
	"log"
	"os"
	"portfoliosite_v4_admin_auth_service/internal/api"
	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found")
    }

	ginMode := os.Getenv("GIN_MODE")
    if ginMode != "" {
        gin.SetMode(ginMode)
    } else {
        log.Println("GIN_MODE not set, defaulting to debug mode")
    }
    jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in the environment")
	}

    gromDB := db.InitDB()
    accessTokenDuration := 15 * time.Minute  

	jwtManager := jwtmanager.NewJWTManager(jwtSecret, accessTokenDuration)
    router := api.SetupRouter(gromDB, jwtManager)
    router.Run() 
}