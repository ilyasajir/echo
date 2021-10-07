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
	users := e.Group("/users")
	users.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	users.GET("", controller.GetUserController)
	m.LogMiddleware(e)

	//Users
	users.GET("/:id", controller.GetUserControllerById)
	users.DELETE("/:id", controller.DeleteUserController)
	users.PUT("/:id", controller.UpdateUserController)
	users.POST("", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	//Pets
	e.GET("/pets", controller.GetPetController)
	e.POST("/pets", controller.CreatePetController)
	e.DELETE("/pets/:id", controller.DeletePetsController)
	e.PUT("/pets/:id", controller.UpdatePetsController)

	//Products
	e.GET("/products", controller.GetProductController)
	e.POST("/products", controller.CreateProductController)
	e.DELETE("/products/:id", controller.DeleteProductController)
	e.PUT("/products/:id", controller.UpdateProductController)

	//Categories
	e.GET("/categories", controller.GetCategoryController)
	e.POST("/categories", controller.CreateCategoryController)
	e.DELETE("/categories/:id", controller.DeleteCategoryController)
	e.PUT("/categories/:id", controller.UpdateCategoryController)

	//Articles
	e.GET("/articles", controller.GetArticleController)
	e.POST("/articles", controller.CreateArticleController)
	e.DELETE("/articles/:id", controller.DeleteArticleController)
	e.PUT("/articles/:id", controller.UpdateArticleController)

	//
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUserController)

	return e

}
