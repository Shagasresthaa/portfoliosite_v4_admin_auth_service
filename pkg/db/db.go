package db

import (
	"fmt"
	"log"
	"os"
	"portfoliosite_v4_admin_auth_service/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found")
    }

    // Construct the data source name from environment variables
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSLMODE"),
    )

    // Open database connection with GORM
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Check database connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get generic database object from GORM DB: %v", err)
    }
    if err = sqlDB.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Connected to the database successfully")

    // Auto migrate the schema
    if err := DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Failed to auto-migrate database schema: %v", err)
    }

    return DB
}

func GetDB() *gorm.DB {
    return DB
}
