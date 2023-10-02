package routes

import (
	"go-elden-ring-api/controllers"

	"github.com/labstack/echo/v4"
)

func RecipeRoutes(e *echo.Echo) {
	e.GET("/api/recipes", controllers.RecipeRead)
	e.GET("/api/recipes/:id", controllers.RecipeShow)
	e.POST("/api/recipes", controllers.RecipeCreate)
	e.PATCH("/api/recipes/:id", controllers.RecipeUpdate)
	e.DELETE("/api/recipes/:id", controllers.RecipeDelete)
	e.POST("/api/recipes/:id/material", controllers.RecipeAddMaterial)
	e.PATCH("/api/recipes/:recipe_id/material/:material_id", controllers.RecipeUpdateMaterial)
	e.DELETE("/api/recipes/:recipe_id/material/:material_id", controllers.RecipeDeleteMaterial)
}
