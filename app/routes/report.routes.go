package routes

import (
	"tampayang-backend/app/controllers"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ReportRoute(app *fiber.App) {
	app.Post("/reports/", controllers.CreateReport)
	app.Get("/reports/urgently-report/", middlewares.AuthCookie, controllers.UrgencyReport)
	app.Get("/reports/manage-report/", middlewares.AuthCookie, controllers.ManageReport)
	app.Get("/reports/detail-report/", middlewares.AuthCookie, controllers.DetailReport)
	app.Put("/reports/update-report/:report_number", middlewares.AuthCookie, controllers.UpdateReport)
}
