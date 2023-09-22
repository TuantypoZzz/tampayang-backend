package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
)

func EmployeeRoute(app *fiber.App) {
	app.Post("/employee/", controllers.CreateEmployeeHandler)
	// app.Get("/user/all-users", user_handler.GetAllUsersHandler)
}
