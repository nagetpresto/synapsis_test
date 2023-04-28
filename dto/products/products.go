package productsdto

type CreateProductRequest struct {
	CategoryID  int    `json:"category_id" form:"category_id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	CategoryID  int    `json:"category_id" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"-" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}