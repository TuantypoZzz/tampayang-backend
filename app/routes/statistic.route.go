package routes

import (
	"tampayang-backend/app/controllers"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func StatisticRoute(app *fiber.App) {
	app.Get("/statistic/report-summary", middlewares.IsLogin, controllers.ReportSummary)
	app.Get("/statistic/report-weekly", controllers.ReportWeekly)
	app.Get("/statistic/report-map", controllers.ReportMap)
}
