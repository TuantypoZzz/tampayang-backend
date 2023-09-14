package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	// User Routes
	UserRoute(app)
}