package main

import (
	. "AOJ-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(recover.New())

	env := LoadEnv()

	app.Get("/api/health", healthCheck)

	log.Fatal(app.Listen(":" + env["BE_PORT"]))
}

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("Backend Alive")
}
