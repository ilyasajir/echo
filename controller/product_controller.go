package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"echo/config"
	"echo/model"

	"github.com/labstack/echo"
)

func GetProductController(c echo.Context) error {
	var products []model.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    products,
	})
}

func CreateProductController(c echo.Context) error {
	//binding data
	product := model.Product{}
	c.Bind(&product)

	err := config.DB.Create(&product).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create product",
		"product":  product,
	})
}

func DeleteProductController(c echo.Context) error {
	var products []model.Product
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := config.DB.Where("product_id=?", id).Delete(&products).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    products,
	})
}

func UpdateProductController(c echo.Context) error {
	var products model.Product
	c.Bind(&products)
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.Where("product_id=?", id).Updates(&products).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    products,
	})
}
