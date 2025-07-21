package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	AuthRoute(app)
	LovRoute(app)
	StatisticRoute(app)
	ReportRoute(app)
}
