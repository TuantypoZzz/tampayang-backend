package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
)

func LovRoute(app *fiber.App) {
	app.Get("/lov/infrastucture-category", controllers.InfrastructureCategory)
	app.Get("/lov/damage-type", controllers.DamageType)
	app.Get("/lov/regency", controllers.Regency)
	app.Get("/lov/district", controllers.District)
	app.Get("/lov/village", controllers.Village)
}
