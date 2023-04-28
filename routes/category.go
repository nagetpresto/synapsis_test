package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"
	"BE/pkg/middleware"

	"github.com/labstack/echo/v4"
)


func CategoryRoutes(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository, cartRepository)

	e.GET("/category", (h.GetAllCategory))
	e.GET("/category/:id", (h.GetOneCategory))
	e.POST("/category", middleware.AdminOnly(middleware.UploadFile(h.CreateCategory)))
	e.PATCH("/category/:id", middleware.AdminOnly(middleware.UploadFile(h.UpdateCategory)))
	e.DELETE("/category/:id", middleware.AdminOnly(h.DeleteCategory))
}