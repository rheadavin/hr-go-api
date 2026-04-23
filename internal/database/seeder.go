package database

import (
	"log"

	"github.com/rheadavin/hr-go-api/internal/models"
)

func Seed() {
	// Cek jika data sudah ada
	var count int64
	DB.Model(&models.Division{}).Count(&count)
	if count > 0 {
		return
	}

	// Seed data
	divisions := []models.Division{
		{Name: "Engineering", Description: "Tim teknologi dan pengembangan"},
		{Name: "HR", Description: "Tim sumber daya manusia"},
		{Name: "Finance", Description: "Tim keuangan dan akuntansi"},
		{Name: "Marketing", Description: "Tim pemasaran dan branding"},
	}

	DB.Create(&divisions)

	log.Println("Seed data created!")
}
