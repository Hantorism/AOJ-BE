package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	backPort := os.Getenv("BE_PORT")

	e.GET("/api/health", healthCheck)

	e.Logger.Fatal(e.Start(":" + backPort))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Backend Alive")
}
