package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	AuthRoute(app)
	TestRoute(app)
	ExampleRoute(app)
	EmployeeRoute(app)
	GolonganRoute(app)
}
