package main

import (
	"go-elden-ring-api/initializers"
	"go-elden-ring-api/routes"

	_ "go-elden-ring-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

// @title Elden Ring API
// @version 1.0
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.MaterialRoutes(e)
	routes.RecipeRoutes(e)
	routes.CookBookRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
