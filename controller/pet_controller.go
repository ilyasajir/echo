package controller

import (
	"fmt"
	"net/http"
	"strconv"

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

	err := config.DB.Create(&pet).Error
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

func DeletePetsController(c echo.Context) error {
	var pets []model.Pet
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := config.DB.Where("pet_id=?", id).Delete(&pets).Error

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
