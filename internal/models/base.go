package models

import (
	"time"

	"gorm.io/gorm"
)

// === BASE MODEL (embed ke semua model) ===
// Seperti BaseEntity di TypeORM / NestJS
type Base struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // soft delete
}
