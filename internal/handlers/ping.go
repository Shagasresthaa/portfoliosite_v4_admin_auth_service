package handlers

import (
	"log"
	"net/http"
	"portfoliosite_v4_admin_auth_service/pkg/db"

	"github.com/gin-gonic/gin"
)

// Just to play ping pong with server lol
// Just a test endpoint take it easy
func PingHandler(c *gin.Context) {
    var tables []string
    result := db.GetDB().Raw(`SELECT tablename FROM pg_tables WHERE schemaname='public'`).Scan(&tables)
    if result.Error != nil {
        log.Println("Failed to execute query: ", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "pong", "tables": tables})
}
