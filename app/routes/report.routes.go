package routes

import (
	"tampayang-backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ReportRoute(app *fiber.App) {
	app.Post("/reports/", controllers.CreateReports)
}
