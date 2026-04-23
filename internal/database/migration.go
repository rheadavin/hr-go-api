package database

import (
	"log"

	"github.com/rheadavin/hr-go-api/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Division{},
		&models.Employee{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed!")
}
