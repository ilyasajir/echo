package route

import (
	"echo/constants"
	"echo/controller"
	"echo/middleware"

	m "echo/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//routing w/ query parameters
	e.GET("/users", controller.GetUserController)
	m.LogMiddleware(e)
	e.GET("/users/:id", controller.GetUserControllerById)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	e.GET("/pets", controller.GetPetController)
	e.POST("/pets", controller.CreatePetController)
	e.DELETE("/pets/:id", controller.DeletePetsController)
	e.PUT("/pets/:id", controller.UpdatePetsController)

	e.GET("/products", controller.GetProductController)
	e.POST("/products", controller.CreateProductController)
	e.DELETE("/products/:id", controller.DeleteProductController)
	e.PUT("/products/:id", controller.UpdateProductController)

	e.GET("/categories", controller.GetCategoryController)
	e.POST("/categories", controller.CreateCategoryController)
	e.DELETE("/categories/:id", controller.DeleteCategoryController)
	e.PUT("/categories/:id", controller.UpdateCategoryController)

	e.GET("/articles", controller.GetArticleController)
	e.POST("/articles", controller.CreateArticleController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUserController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUserController)
	return e

}
