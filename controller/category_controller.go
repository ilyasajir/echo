package controller

import (
	"fmt"
	"net/http"
	"strconv"

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

	err := config.DB.Create(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success adding category",
		"category": category,
	})
}

func DeleteCategoryController(c echo.Context) error {
	var categories []model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := config.DB.Where("category_id=?", id).Delete(&categories).Error

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
