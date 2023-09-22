package routes

import (
	"github.com/gofiber/fiber/v2"
	employeecontroller_handler "github.com/nulla-vis/golang-fiber-template/app/handler/employee_controller"
)

func EmployeeRoute(app *fiber.App) {
	app.Post("/employee/", employeecontroller_handler.CreateEmployeeHandler)
	// app.Get("/user/all-users", user_handler.GetAllUsersHandler)
}
