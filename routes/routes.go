package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/auth-jwt-golang/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/user", controllers.User)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}
