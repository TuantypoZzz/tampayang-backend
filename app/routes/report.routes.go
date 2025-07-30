package routes

import (
	"tampayang-backend/app/controllers"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ReportRoute(app *fiber.App) {
	app.Post("/reports/", controllers.CreateReport)
	app.Get("/reports/check-status/", controllers.CheckStatus)
	app.Get("/reports/urgently/", middlewares.IsLogin, controllers.UrgencyReport)
	app.Get("/reports/manage/", middlewares.AuthCookie, controllers.ManageReport)
	app.Get("/reports/:report_id", middlewares.AuthCookie, controllers.DetailReport)
	app.Put("/reports/:report_id", middlewares.AuthCookie, controllers.UpdateReport)
}
