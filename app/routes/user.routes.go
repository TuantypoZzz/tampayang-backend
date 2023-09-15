package routes

import (
	"github.com/gofiber/fiber/v2"
	user_handler "github.com/nulla-vis/golang-fiber-template/app/handler/user"
)

func UserRoute(app *fiber.App) {
	// User Routes
	// app.Get("/user/:name", handler.CreateUserHandler)
	app.Get("/user/all-user", user_handler.GetAllUserHandler)
}