package signIn

import (
	. "AOJ-BE/models"
	. "AOJ-BE/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func SignIn(c *fiber.Ctx) error {
	body := new(struct {
		Email    string
		Password string
	})
	err := c.BodyParser(body)
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	var user User
	query := "SELECT email, nickname, student_number, password FROM \"user\" WHERE email = $1"
	err = DB.QueryRow(query, body.Email).Scan(&user.Email, &user.Nickname, &user.StudentNumber, &user.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"result": "user not found",
			})
		}
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	if body.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"result": "password not match",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
