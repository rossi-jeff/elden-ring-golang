package controllers

import (
	"go-elden-ring-api/initializers"
	"go-elden-ring-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MaterialRead List Materials Alphabetically
//
// @Summary Get Materials
// @Description List Materials Alphabetically
// @Tags         Materials
// @Accept       json
// @Produce      json
//
//	@Success      200  {array}   models.Material
//
// @Router /api/materials [get]
func MaterialRead(c echo.Context) error {
	materials := []models.Material{}
	initializers.DB.Order("Name asc").Find(&materials)
	return c.JSON(http.StatusOK, materials)
}

// MaterialShow List Materials Alphabetically
//
// @Summary Get a Material
// @Description Get a Material by ID
// @Tags         Materials
// @Accept       json
// @Produce      json
// @Param MaterialId path int true "Material ID"
//
//	@Success      200  {array}   models.Material
//
// @Router /api/materials/{MaterialId} [get]
func MaterialShow(c echo.Context) error {
	material := models.Material{}
	initializers.DB.Where("ID = ?", c.Param("id")).First(&material)
	return c.JSON(http.StatusOK, material)
}
