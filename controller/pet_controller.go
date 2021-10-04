package controller

import (
	"net/http"

	"echo/config"
	"echo/model"

	"github.com/labstack/echo"
)

func GetPetController(c echo.Context) error {
	var pets []model.Pet

	err := config.DB.Find(&pets).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    pets,
	})
}

func CreatePetController(c echo.Context) error {
	//binding data
	pet := model.Pet{}
	c.Bind(&pet)

	err := config.DB.Save(&pet).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create pet",
		"pet":      pet,
	})
}
