package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/joho/godotenv"
)

var (
	DB  *sql.DB
	ENV = make(map[string]string)

	HealthCheckConfig = healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/apis/livez",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/apis/readyz",
	}
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ENV["BE_PORT"] = os.Getenv("BE_PORT")
	ENV["DB_HOST"] = os.Getenv("DB_HOST")
	ENV["DB_PORT"] = os.Getenv("DB_PORT")
	ENV["DB_USER"] = os.Getenv("DB_USER")
	ENV["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	ENV["DB_NAME"] = os.Getenv("DB_NAME")
}

func ConnectDB() {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		ENV["DB_HOST"], ENV["DB_PORT"], ENV["DB_USER"], ENV["DB_PASSWORD"], ENV["DB_NAME"])

	var err error
	DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
