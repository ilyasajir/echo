package controller

import (
	"net/http"

	"echo/config"
	"echo/model"

	"github.com/labstack/echo"
)

func GetArticleController(c echo.Context) error {
	var articles []model.Article

	err := config.DB.Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    articles,
	})
}

func CreateArticleController(c echo.Context) error {
	//binding data
	article := model.Article{}
	c.Bind(&article)

	err := config.DB.Create(&article).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create article",
		"article":  article,
	})
}
