package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	ProductRoutes(e)
	CartRoutes(e)
	TransactionRoutes(e)
	AuthRoutes(e)
}