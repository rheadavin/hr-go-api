package dto

type CreateDivisionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateDivisionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ListDivisionRequest struct {
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1"`
	Search string `form:"search"`
}

type DivisionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
