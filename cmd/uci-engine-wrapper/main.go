package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/calculate", handlers.Calculate)

	e.Logger.Fatal(e.Start(":1323"))
}
