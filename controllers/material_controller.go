package controllers

import (
	"go-elden-ring-api/dto"
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
//	@Success      200  {object}   models.Material
//
// @Router /api/materials/{MaterialId} [get]
func MaterialShow(c echo.Context) error {
	material := models.Material{}
	initializers.DB.Where("ID = ?", c.Param("id")).First(&material)
	return c.JSON(http.StatusOK, material)
}

// MaterialCreate Create a New Material
//
// @Summary Create Material
// @Description Create a New Material
// @Tags         Materials
// @Accept       json
// @Produce      json
// @Param	data	body	dto.NameDescription	true	"New Material Params"
//
//	@Success      201  {object}   models.Material
//
// @Router /api/materials [post]
func MaterialCreate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	material := models.Material{Name: body.Name, Description: body.Description}
	initializers.DB.Create(&material)
	return c.JSON(http.StatusCreated, material)
}

// MaterialUpdate Update a Material by ID
//
// @Summary Update a Material
// @Description Update a Material by ID
// @Tags         Materials
// @Accept       json
// @Produce      json
// @Param MaterialId path int true "Material ID"
// @Param	data	body	dto.NameDescription	true	"Update Material Params"
//
//	@Success      200  {object}   models.Material
//
// @Router /api/materials/{MaterialId} [patch]
func MaterialUpdate(c echo.Context) error {
	body := dto.NameDescription{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	material := models.Material{}
	initializers.DB.First(&material, c.Param("id"))
	material.Name = body.Name
	material.Description = body.Description
	initializers.DB.Save(&material)
	return c.JSON(http.StatusOK, material)
}

// MaterialDelete Delete a Material by ID
//
// @Summary Delete a Material
// @Description Delete a Material by ID
// @Tags         Materials
// @Accept       json
// @Produce      json
// @Param MaterialId path int true "Material ID"
//
//	@Success      204
//
// @Router /api/materials/{MaterialId} [delete]
func MaterialDelete(c echo.Context) error {
	initializers.DB.Delete(&models.Material{}, c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}
