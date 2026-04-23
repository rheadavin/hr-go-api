package models

import (
	"github.com/rheadavin/hr-go-api/pkg/types"
)

type Employee struct {
	Base
	NIK        string     `json:"nik" gorm:"not null;uniqueIndex;size:20"`
	FullName   string     `json:"full_name" gorm:"not null;size:150"`
	Email      string     `json:"email" gorm:"not null;uniqueIndex;size:100"`
	Phone      string     `json:"phone" gorm:"size:20"`
	Position   string     `json:"position" gorm:"size:100"`
	Salary     float64    `json:"salary" gorm:"type:decimal(15,2)"`
	JoinDate   types.Date `json:"join_date" gorm:"type:date"`
	IsActive   bool       `json:"is_active" gorm:"default:true"`
	DivisionID uint       `json:"division_id" gorm:"not null;index"`
	Division   Division   `json:"division,omitempty" gorm:"foreignKey:DivisionID"`
}

func (Employee) TableName() string { return "employees" }
