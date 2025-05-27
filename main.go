package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/short", handleShortJSON)
	e.GET("/long", handleLongJSON)

	e.Logger.Fatal(e.Start(":8080"))
}
