package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/health", healthCheck)

	e.Logger.Fatal(e.Start(":1323"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Backend Alive")
}
