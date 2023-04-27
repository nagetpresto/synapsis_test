package productsdto

type CreateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}