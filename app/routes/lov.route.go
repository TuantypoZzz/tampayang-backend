package routes

import (
	"tampayang-backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func LovRoute(app *fiber.App) {
	app.Get("/lov/infrastructure-category", controllers.InfrastructureCategory)
	app.Get("/lov/damage-type", controllers.DamageType)
	app.Get("/lov/province", controllers.Province)
	app.Get("/lov/regency", controllers.Regency)
	app.Get("/lov/district", controllers.District)
	app.Get("/lov/village", controllers.Village)
}
