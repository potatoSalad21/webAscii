package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/potatoSalad21/webAscii/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/static/js", "static/js")

	e.GET("/", handlers.HandleHome)

	e.Logger.Fatal(e.Start(":8080"))
}
