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
	app.Get("/statistic/dashboard-summary", controllers.DashboardSummary)

	// New advanced analytics endpoints
	app.Get("/statistic/monthly-report", controllers.MonthlyReport)
	app.Get("/statistic/category-breakdown", controllers.CategoryBreakdown)
	app.Get("/statistic/performance-chart", controllers.PerformanceChart)
}
