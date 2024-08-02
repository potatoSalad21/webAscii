package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/potatoSalad21/webAscii/views/home"
)

func handleHome(echo.Context) error {
    return nil
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.GET("/", handleHome)

    e.Logger.Fatal(e.Start(":8080"))
}
