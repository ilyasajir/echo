package controller

import (
	"fmt"
	"net/http"
	"strconv"

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

func DeleteArticleController(c echo.Context) error {
	var articles []model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := config.DB.Where("category_id=?", id).Delete(&articles).Error

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

func UpdateArticleController(c echo.Context) error {
	var article model.Article
	c.Bind(&article)
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.Where("article_id=?", id).Updates(&article).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    article,
	})
}
