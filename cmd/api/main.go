package main

import (
	"fmt"
	"log"

	"github.com/rheadavin/hr-go-api/internal/config"
	"github.com/rheadavin/hr-go-api/internal/database"
	"github.com/rheadavin/hr-go-api/router"
)

func main() {
	// 1. Load konfigurasi
	config.Load()

	// 2. init database
	database.Connect()

	// 3. run migration
	database.Migrate()

	// 4. seed data
	database.Seed()

	// 5. setup router
	r := router.SetupRouter(database.DB)

	// 6. start server
	addr := fmt.Sprintf(":%s", config.AppConfig.AppPort)
	log.Printf("Server running on http://localhost%s", addr)
	log.Fatal(r.Run(addr))
}
