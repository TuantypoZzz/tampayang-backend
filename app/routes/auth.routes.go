package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
	"github.com/nulla-vis/golang-fiber-template/core/middlewares"
)

func AuthRoute(app *fiber.App) {
	app.Post("/auth/login/", middlewares.IsLogin, controllers.Login)
	app.Post("/auth/logout/", middlewares.IsLogin, controllers.Logout)
}