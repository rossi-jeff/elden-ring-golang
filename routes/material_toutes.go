package routes

import (
	"go-elden-ring-api/controllers"

	"github.com/labstack/echo/v4"
)

func MaterialRoutes(e *echo.Echo) {
	e.GET("/api/materials", controllers.MaterialRead)
	e.GET("/api/materials/:id", controllers.MaterialShow)
	e.POST("/api/materials", controllers.MaterialCreate)
	e.PATCH("/api/materials/:id", controllers.MaterialUpdate)
	e.DELETE("/api/materials/:id", controllers.MaterialDelete)
}
