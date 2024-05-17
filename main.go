package main

import (
	. "AOJ-BE/src/apis"
	. "AOJ-BE/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(healthcheck.New(HealthCheckConfig))

	LoadENV()

	ConnectDB()
	defer DB.Close()

	app.Get("/api/health", HealthCheck)
	app.Post("/api/signin", Signin)

	log.Fatal(app.Listen(":" + ENV["BE_PORT"]))
}
