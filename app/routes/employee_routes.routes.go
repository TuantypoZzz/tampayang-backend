package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
	"github.com/nulla-vis/golang-fiber-template/core/middlewares"
)

func EmployeeRoute(app *fiber.App) {
	app.Post("/employee/", middlewares.AuthCookie, controllers.CreateEmployeeHandler)
	app.Get("/employee/all-employee", middlewares.AuthCookie, controllers.GetAllEmployee)
	app.Get("/employee/get-employee/:employee_id", middlewares.AuthCookie, controllers.GetEmployeeById)
}
