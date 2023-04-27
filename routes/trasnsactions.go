package routes

import (
	"BE/handlers"
	"BE/pkg/middleware"
	"BE/pkg/mysql"
	"BE/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transactions", h.GetAllTransaction)
	e.GET("/transactions/:id", h.GetOneTransaction)
	
	e.GET("/transactions-user", middleware.Auth(h.GetUserTrans))
	e.PATCH("/transactions", middleware.Auth(h.DoTransaction))
	
	e.PATCH("/transactions/:id", (middleware.AdminOnly(h.UpdateTransaction)))
	e.DELETE("/transactions/:id", middleware.Auth(h.DeleteTransaction))
	e.POST("/notification", h.Notification)
}