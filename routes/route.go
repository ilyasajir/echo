package route

import (
	"echo/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	//routing w/ query parameters
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/pets", controller.GetPetController)
	e.POST("/pets", controller.CreatePetController)
	return e

}
