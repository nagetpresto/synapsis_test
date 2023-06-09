package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"
	"BE/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	CartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(CartRepository)

	e.GET("/carts",  h.GetAllCart)
	e.GET("/carts/:id",  h.GetOneCart)
	e.GET("/carts-active", middleware.Auth(h.GetActiveCart))
	e.POST("/carts",  middleware.Auth(h.CreateCart))
	e.PATCH("/carts/:id",  middleware.Auth(h.UpdateCart))
	e.DELETE("/carts/:id",  middleware.Auth(h.DeleteCart))
}