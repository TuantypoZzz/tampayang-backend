package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/handler/test_controller"
)

func TestRoute(app *fiber.App) {
	// User Routes
	// app.Get("/user/:name", test_controller.CreateUserHandler)
	app.Get("/user/all-user", test_controller.GetAllUserHandler)
}