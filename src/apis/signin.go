package apis

import (
	. "AOJ-BE/models"
	. "AOJ-BE/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Signin(c *fiber.Ctx) error {
	signin := new(SignIn)
	err := c.BodyParser(signin)
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}

	rows, err := DB.Query("SELECT * FROM \"user\"")
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": "error on server",
		})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.Nickname, &user.StudentNumber, &user.Password)
		if err != nil {
			log.Panic(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"result": "error on server",
			})
		}
		users = append(users, user)
	}

	user, exist := contains(users, signin.Email)
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"result": "user not found",
		})
	}

	if signin.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"result": "password not match",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func contains(s []User, target string) (User, bool) {
	for _, v := range s {
		if v.Email == target {
			return v, true
		}
	}
	return User{}, false
}
