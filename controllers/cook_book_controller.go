package controllers

import (
	"go-elden-ring-api/dto"
	"go-elden-ring-api/initializers"
	"go-elden-ring-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CookBookRead lists all Cook Books
//
// @Summary Get Cook Books
// @Description List Cook Books Alphabetically
// @Tags         Cook Books
// @Accept       json
// @Produce      json
//
//	@Success      200  {array}   models.CookBook
//
// @Router /api/cook_books [get]
func CookBookRead(c echo.Context) error {
	cookBooks := []models.CookBook{}
	initializers.DB.Order("Name asc").Find(&cookBooks)
	return c.JSON(http.StatusOK, cookBooks)
}

// CookBookShow Finds a Cook Book by ID
//
// @Summary Get a Cook Book
// @Description Finds a Cook Book by ID
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param CookBookID path int true "Cook Book ID"
//
//	@Success      200  {object}   models.CookBook
//
// @Router /api/cook_books/{CookBookID} [get]
func CookBookShow(c echo.Context) error {
	cookBook := models.CookBook{}
	initializers.DB.Where("ID = ?", c.Param("id")).Preload("Recipes").First(&cookBook)
	return c.JSON(http.StatusOK, cookBook)
}

// @Summary Create a Cook Book
// @Description Create a Cook Book
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param	data	body	dto.NameDescription	true	"New Cook Book Params"
//
//	@Success      201  {object}   models.CookBook
//
// @Router /api/cook_books [post]
func CookBookCreate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookBook := models.CookBook{Name: body.Name, Description: body.Description}
	initializers.DB.Create(&cookBook)
	return c.JSON(http.StatusCreated, cookBook)
}

// @Summary Update a Cook Book
// @Description Update a Cook Book by ID
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param CookBookID path int true "Cook Book ID"
// @Param	data	body	dto.NameDescription	true	"Update Cook Book Params"
//
//	@Success      200  {object}   models.CookBook
//
// @Router /api/cook_books/{CookBookID} [patch]
func CookBookUpdate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookBook := models.CookBook{}
	initializers.DB.First(&cookBook, c.Param("id"))
	cookBook.Name = body.Name
	cookBook.Description = body.Description
	initializers.DB.Save(&cookBook)
	return c.JSON(http.StatusOK, cookBook)
}

// @Summary Delete a Cook Book
// @Description Delete a Cook Book by ID
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param CookBookID path int true "Cook Book ID"
//
//	@Success      204
//
// @Router /api/cook_books/{CookBookID} [delete]
func CookBookDelete(c echo.Context) error {
	initializers.DB.Delete(&models.CookBook{}, c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}

// @Summary Create a Cook Book Recipe
// @Description Adds a Recipe to Cook Book by ID
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param CookBookID path int true "Cook Book ID"
// @Param RecipeID path int true "Recipe ID"
//
//	@Success      201  {object}   models.CookBookRecipe
//
// @Router /api/cook_books/{CookBookID}/recipe/{RecipeID} [post]
func CookBookAddRecipe(c echo.Context) error {
	cookBook := models.CookBook{}
	initializers.DB.First(&cookBook, c.Param("cook_book_id"))
	recipe := models.Recipe{}
	initializers.DB.First(&recipe, c.Param("recipe_id"))
	cookBookRecipe := models.CookBookRecipe{CookBookID: cookBook.ID, RecipeID: recipe.ID}
	initializers.DB.Create(&cookBookRecipe)
	return c.JSON(http.StatusCreated, cookBookRecipe)
}

// @Summary Delete a Cook Book Recipe
// @Description Delete a Recipe from Cook Book
// @Tags         Cook Books
// @Accept       json
// @Produce      json
// @Param CookBookID path int true "Cook Book ID"
// @Param RecipeID path int true "Recipe ID"
//
//	@Success      204
//
// @Router /api/cook_books/{CookBookID}/recipe/{RecipeID} [delete]
func CookBookDeleteRecipe(c echo.Context) error {
	initializers.DB.Where("CookBookID = ? and RecipeID = ?", c.Param("cook_book_id"), c.Param("recipe_id")).Delete(&models.CookBookRecipe{})
	return c.NoContent(http.StatusNoContent)
}
