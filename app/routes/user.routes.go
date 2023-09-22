package routes

import (
	"github.com/gofiber/fiber/v2"
	user_handler "github.com/nulla-vis/golang-fiber-template/app/handler/test_controller/user"
)

func UserRoute(app *fiber.App) {
	app.Post("/user/", user_handler.CreateUserHandler)
	app.Get("/user/all-users", user_handler.GetAllUsersHandler)
}
