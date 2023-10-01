package controllers

import (
	"go-elden-ring-api/initializers"
	"go-elden-ring-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RecipeRead List Recipes Alphabetically
//
// @Summary Get Recipes
// @Description List Recipes Alphabetically
// @Tags         Recipes
// @Accept       json
// @Produce      json
//
//	@Success      200  {array}   models.Recipe
//
// @Router /api/recipes [get]
func RecipeRead(c echo.Context) error {
	recipes := []models.Recipe{}
	initializers.DB.Order("Name asc").Find(&recipes)
	return c.JSON(http.StatusOK, recipes)
}

// RecipeShow Get a Recipe by ID
//
// @Summary Get a Recipe
// @Description Get a Recipe by ID
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param RecipeID path int true "Recipe ID"
//
//	@Success      200  {array}   models.Recipe
//
// @Router /api/recipes/{RecipeID} [get]
func RecipeShow(c echo.Context) error {
	recipe := models.Recipe{}
	initializers.DB.Where("ID = ?", c.Param("id")).Preload("Materials").Preload("CookBooks").Preload("Materials.Material").First(&recipe)
	return c.JSON(http.StatusOK, recipe)
}
