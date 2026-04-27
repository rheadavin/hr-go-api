package dto

import (
	"github.com/rheadavin/hr-go-api/pkg/types"
)

type CreateEmployeeRequest struct {
	NIK        string     `json:"nik" binding:"required"`
	FullName   string     `json:"full_name" binding:"required"`
	Email      string     `json:"email" binding:"required,email"`
	Phone      string     `json:"phone" binding:"required"`
	Position   string     `json:"position" binding:"required"`
	Salary     float64    `json:"salary" binding:"required"`
	JoinDate   types.Date `json:"join_date" binding:"required" time_format:"2006-01-02"`
	DivisionID uint       `json:"division_id" binding:"required"`
}

type UpdateEmployeeRequest struct {
	NIK        string     `json:"nik" binding:"required"`
	FullName   string     `json:"full_name" binding:"required"`
	Email      string     `json:"email" binding:"required,email"`
	Phone      string     `json:"phone" binding:"required"`
	Position   string     `json:"position" binding:"required"`
	Salary     float64    `json:"salary" binding:"required"`
	JoinDate   types.Date `json:"join_date" binding:"required" time_format:"2006-01-02"`
	DivisionID uint       `json:"division_id" binding:"required"`
}

type ListEmployeeRequest struct {
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1"`
	Search string `form:"search"`
}

type EmployeeResponse struct {
	ID         uint             `json:"id"`
	NIK        string           `json:"nik"`
	FullName   string           `json:"full_name"`
	Email      string           `json:"email"`
	DivisionID uint             `json:"division_id"`
	Division   DivisionResponse `json:"division"`
}
