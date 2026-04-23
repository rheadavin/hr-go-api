package models

type Division struct {
	Base
	Name        string     `json:"name" gorm:"not null;uniqueIndex;size:100"`
	Description string     `json:"description" gorm:"size:500"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
	Employees   []Employee `json:"employees,omitempty" gorm:"foreignKey:DivisionID"`
}

func (Division) TableName() string { return "divisions" }
