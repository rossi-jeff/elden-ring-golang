package routes

import (
	"go-elden-ring-api/controllers"

	"github.com/labstack/echo/v4"
)

func CookBookRoutes(e *echo.Echo) {
	e.GET("/api/cook_books", controllers.CookBookRead)
	e.GET("/api/cook_books/:id", controllers.CookBookShow)
}
