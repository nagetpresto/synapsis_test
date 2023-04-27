package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"
	"BE/pkg/middleware"

	"github.com/labstack/echo/v4"
)


func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", (h.GetAllProduct))
	e.GET("/products/:id", (h.GetOneProduct))
	e.POST("/products", middleware.AdminOnly(middleware.UploadFile(h.CreateProduct)))
	e.PATCH("/products/:id", middleware.AdminOnly(middleware.UploadFile(h.UpdateProduct)))
	e.DELETE("/products/:id", middleware.AdminOnly(h.DeleteProduct))
}