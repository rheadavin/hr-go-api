package main

import (
	"fmt"
	"log"

	_ "github.com/rheadavin/hr-go-api/docs"
	"github.com/rheadavin/hr-go-api/internal/config"
	"github.com/rheadavin/hr-go-api/internal/database"
	"github.com/rheadavin/hr-go-api/router"
)

// @title MyAPI - Sistem Manajemen Karyawan
// @version 1.0
// @description REST API untuk manajemen divisi dan karyawan
// @termsOfService http://localhost:8080.com/terms/
// @contact.name API Support
// @contact.email rheadavin13@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: Bearer {token}

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
