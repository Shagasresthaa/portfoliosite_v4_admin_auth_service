package main

import (
	"portfoliosite_v4_admin_auth_service/internal/api"
	"portfoliosite_v4_admin_auth_service/pkg/db"
)

func main() {
    gromDB := db.InitDB()
    router := api.SetupRouter(gromDB)
    router.Run() // Default runs on PORT 8080
}