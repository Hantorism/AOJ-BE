package main

import (
	. "AOJ-BE/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	env := LoadEnv()

	e.GET("/api/health", healthCheck)

	e.Logger.Fatal(e.Start(":" + env["backPort"]))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Backend Alive")
}
