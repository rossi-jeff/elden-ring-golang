package controllers

import (
	"go-elden-ring-api/dto"
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
//	@Success      200  {object}   models.Recipe
//
// @Router /api/recipes/{RecipeID} [get]
func RecipeShow(c echo.Context) error {
	recipe := models.Recipe{}
	initializers.DB.Where("ID = ?", c.Param("id")).Preload("Materials").Preload("CookBooks").Preload("Materials.Material").First(&recipe)
	return c.JSON(http.StatusOK, recipe)
}

// @Summary Create a Recipe
// @Description Create a Recipe
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param	data	body	dto.NameDescription	true	"New Recipe Params"
//
//	@Success      201  {object}   models.Recipe
//
// @Router /api/recipes [post]
func RecipeCreate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	recipe := models.Recipe{Name: body.Name, Description: body.Description}
	initializers.DB.Create(&recipe)
	return c.JSON(http.StatusCreated, recipe)
}

// @Summary Update a Recipe
// @Description Update a Recipe by ID
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param RecipeID path int true "Recipe ID"
// @Param	data	body	dto.NameDescription	true	"New Recipe Params"
//
//	@Success      200  {object}   models.Recipe
//
// @Router /api/recipes/{RecipeID} [patch]
func RecipeUpdate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	recipe := models.Recipe{}
	initializers.DB.First(&recipe, c.Param("id"))
	recipe.Name = body.Name
	recipe.Description = body.Description
	initializers.DB.Save(&recipe)
	return c.JSON(http.StatusOK, recipe)
}

// @Summary Delete a Recipe
// @Description Delete a Recipe by ID
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param RecipeID path int true "Recipe ID"
//
//	@Success      204
//
// @Router /api/recipes/{RecipeID} [delete]
func RecipeDelete(c echo.Context) error {
	initializers.DB.Delete(&models.Recipe{}, c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}

// @Summary Create a Recipe Material
// @Description Create a Recipe Material
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param	data	body	dto.QuantityMaterialId	true	"New Recipe Material Params"
// @Param RecipeID path int true "Recipe ID"
//
//	@Success      201  {object}   models.RecipeMaterial
//
// @Router /api/recipes/{RecipeID}/material [post]
func RecipeAddMaterial(c echo.Context) error {
	body := dto.QuantityMaterialId{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	recipe := models.Recipe{}
	initializers.DB.First(&recipe, c.Param("id"))
	recipeMaterial := models.RecipeMaterial{RecipeID: recipe.ID, MaterialID: body.MaterialID, Quantity: body.Quantity}
	initializers.DB.Create(&recipeMaterial)
	return c.JSON(http.StatusCreated, recipeMaterial)
}

// @Summary Update a Recipe Material
// @Description Update a Recipe Material
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param	data	body	dto.Quantity	true	"Update Recipe Material Params"
// @Param RecipeID path int true "Recipe ID"
// @Param MaterialID path int true "Material ID"
//
//	@Success      200  {object}   models.RecipeMaterial
//
// @Router /api/recipes/{RecipeID}/material/{MaterialID} [patch]
func RecipeUpdateMaterial(c echo.Context) error {
	body := dto.Quantity{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	recipeMaterial := models.RecipeMaterial{}
	initializers.DB.Where("RecipeId = ? and MaterialId = ?", c.Param("recipe_id"), c.Param("material_id")).First(&recipeMaterial)
	recipeMaterial.Quantity = body.Quantity
	initializers.DB.Save(&recipeMaterial)
	return c.JSON(http.StatusCreated, recipeMaterial)
}

// @Summary Delete a Recipe Material
// @Description Delete a Recipe Material
// @Tags         Recipes
// @Accept       json
// @Produce      json
// @Param RecipeID path int true "Recipe ID"
// @Param MaterialID path int true "Material ID"
//
//	@Success      204
//
// @Router /api/recipes/{RecipeID}/material/{MaterialID} [delete]
func RecipeDeleteMaterial(c echo.Context) error {
	recipeMaterial := models.RecipeMaterial{}
	initializers.DB.Where("RecipeId = ? and MaterialId = ?", c.Param("recipe_id"), c.Param("material_id")).Delete(&recipeMaterial)
	return c.NoContent(http.StatusNoContent)
}
