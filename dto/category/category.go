package categrysdto

type CreateCategoryRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Image       string `json:"image" form:"image"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" form:"name"`
	Image       string `json:"image" form:"image"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Image       string `json:"image" form:"image"`
}