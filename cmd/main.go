package main

import (
	"log"
	"os"
	"portfoliosite_v4_admin_auth_service/internal/api"
	"portfoliosite_v4_admin_auth_service/internal/jwtmanager"
	"portfoliosite_v4_admin_auth_service/pkg/db"
	"time"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found")
    }

    gromDB := db.InitDB()
    jwtManager := jwtmanager.NewJWTManager(os.Getenv("JWT_SECRET"), 72*time.Hour)
    router := api.SetupRouter(gromDB, jwtManager)
    router.Run() // Default runs on PORT 8080
}