package routes

import (
	"go-elden-ring-api/controllers"

	"github.com/labstack/echo/v4"
)

func RecipeRoutes(e *echo.Echo) {
	e.GET("/api/recipes", controllers.RecipeRead)
	e.GET("/api/recipes/:id", controllers.RecipeShow)
}
