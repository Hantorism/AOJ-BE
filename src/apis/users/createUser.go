package users

import (
	. "AOJ-BE/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CreateUser(c *fiber.Ctx) error {
	body := new(struct {
		Email         string `json:"email"`
		Nickname      string `json:"nickname"`
		StudentNumber string `json:"studentNumber"`
		Password      string `json:"password"`
	})
	err := c.BodyParser(body)
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	var email string
	query := "SELECT email FROM \"user\" WHERE email = $1"
	err = DB.QueryRow(query, body.Email).Scan(&email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	if email == body.Email {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result": "user already exist",
		})
	}

	query = "INSERT INTO \"user\" (email, nickname, student_number, password) VALUES ($1, $2, $3, $4)"
	_, err = DB.Exec(query, body.Email, body.Nickname, body.StudentNumber, body.Password)
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "signUp success",
	})
}
