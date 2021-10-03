package main

import (
	"echo/config"
	route "echo/routes"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
