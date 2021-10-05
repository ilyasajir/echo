package route

import (
	"echo/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	//routing w/ query parameters
	e.GET("/users", controller.GetUserController)
	e.GET("/users/:id", controller.GetUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/pets", controller.GetPetController)
	e.POST("/pets", controller.CreatePetController)
	e.DELETE("/pets/:id", controller.DeletePetsController)
	e.GET("/products", controller.GetProductController)
	e.POST("/products", controller.CreateProductController)
	e.DELETE("/products/:id", controller.DeleteProductController)
	e.GET("/categories", controller.GetCategoryController)
	e.POST("/categories", controller.CreateCategoryController)
	e.DELETE("/categories/:id", controller.DeleteCategoryController)
	e.GET("/articles", controller.GetArticleController)
	e.POST("/articles", controller.CreateArticleController)
	return e

}
