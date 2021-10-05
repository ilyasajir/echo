package controller

import (
	"net/http"
	"strconv"

	"echo/config"
	"echo/model"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	var users []model.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.Where("id=?", id).Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	//binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func DeleteUserController(c echo.Context) error {
	var users []model.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.Where("id=?", id).Delete(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
