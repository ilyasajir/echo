package controller

import (
	"net/http"
	"strconv"

	"echo/config"
	"echo/middleware"
	"echo/model"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	var users []model.User
	err := config.DB.Find(&users).Error
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

func GetUserControllerById(c echo.Context) error {
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

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	UserResponse := model.UsersResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"messages": "success create user",
		"user":     UserResponse,
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

func UpdateUserController(c echo.Context) error {
	var users model.User
	c.Bind(&users)
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.Where("id=?", id).Updates(&users).Error

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
