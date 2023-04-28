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
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerProduct(productRepository, cartRepository)

	e.GET("/products/all", (h.GetAllProduct))
	e.GET("/products", (h.GetAllProductbyCategory))
	e.GET("/products/:id", (h.GetOneProduct))
	e.POST("/products", middleware.AdminOnly(middleware.UploadFile(h.CreateProduct)))
	e.PATCH("/products/:id", middleware.AdminOnly(middleware.UploadFile(h.UpdateProduct)))
	e.DELETE("/products/:id", middleware.AdminOnly(h.DeleteProduct))
}