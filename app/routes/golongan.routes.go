package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
	"github.com/nulla-vis/golang-fiber-template/core/middlewares"
)

func GolonganRoute(app *fiber.App) {
	app.Post("/golongan/", middlewares.AuthCookie, controllers.CreateGolongan)
}
