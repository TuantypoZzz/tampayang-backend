package routes

import (
	"tampayang-backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExportRoute(app *fiber.App) {
	// Export endpoints with authentication
	app.Get("/export/reports", controllers.ExportReports)
	app.Get("/export/statistics", controllers.ExportStatistics)
}
