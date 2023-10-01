package controllers

import (
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
