package models

type User struct {
	Base
	Name     string `json:"name" gorm:"not null;size:100"`
	Email    string `json:"email" gorm:"not null;uniqueIndex;size:100"`
	Password string `json:"-" gorm:"not null"` // json:"-" = tidak di-expose ke response
	Role     string `json:"role" gorm:"default:'staff';size:20"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

func (User) TableName() string { return "users" }
