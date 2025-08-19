package routes

import (
	"tampayang-backend/app/controllers"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func MasterRoute(app *fiber.App) {
	// Master data management endpoints with authentication

	// Additional utility endpoints (must come before parameterized routes)
	app.Get("/master/locations/stats", middlewares.AuthCookie, controllers.GetLocationStats)

	// Location CRUD endpoints
	app.Post("/master/locations", middlewares.AuthCookie, controllers.CreateLocation)
	app.Get("/master/locations", middlewares.AuthCookie, controllers.GetLocations)
	app.Get("/master/locations/:id", middlewares.AuthCookie, controllers.GetLocation)
	app.Put("/master/locations/:id", middlewares.AuthCookie, controllers.UpdateLocation)
	app.Delete("/master/locations/:id", middlewares.AuthCookie, controllers.DeleteLocation)

	// Parameterized utility endpoints
	app.Get("/master/locations/:id/dependencies", middlewares.AuthCookie, controllers.CheckLocationDependencies)
}
