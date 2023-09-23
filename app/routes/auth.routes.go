package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
)

func AuthRoute(app *fiber.App) {
	app.Post("/auth/login/", controllers.Login)
}