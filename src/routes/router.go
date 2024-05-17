package routes

import (
	. "AOJ-BE/src/apis/healthCheck"
	. "AOJ-BE/src/apis/signIn"
	. "AOJ-BE/src/apis/users"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", HealthCheck)

	api.Post("/signIn", SignIn)
	//api.Post("/signOut", SignOut)

	users := api.Group("/users")
	users.Post("", CreateUser)
	users.Patch("", UpdateUser)
	users.Delete("", DeleteUser)
}
