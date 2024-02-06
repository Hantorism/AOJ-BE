package main

import (
	. "AOJ-BE/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	//app := fiber.New()
	//
	//app.Use(healthcheck.New())
	//app.Use(logger.New())
	//app.Use(recover.New())
	//
	//env := LoadEnv()
	//
	//app.Get("/api/health", healthCheck)
	//
	//log.Fatal(app.Listen(":" + env["BE_PORT"]))

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	env := LoadEnv()

	e.GET("/api/health", healthCheck)

	e.Logger.Fatal(e.Start(":" + env["BE_PORT"]))
}

//func healthCheck(c *fiber.Ctx) error {
//  return c.Status(fiber.StatusOK).JSON(fiber.Map{
//    "msg": "도커 잘~ 돌아간다",
//  })
//}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Backend Alive")
}
