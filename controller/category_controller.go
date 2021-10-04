package controller

import (
	"net/http"

	"echo/config"
	"echo/model"

	"github.com/labstack/echo"
)

func GetCategoryController(c echo.Context) error {
	var categories []model.Category

	err := config.DB.Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    categories,
	})
}

func CreateCategoryController(c echo.Context) error {
	//binding data
	category := model.Category{}
	c.Bind(&category)

	err := config.DB.Save(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success adding category",
		"pet":      category,
	})
}
