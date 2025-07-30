package routes

import (
	"tampayang-backend/app/controllers"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Post("/auth/login/", middlewares.IsLogin, controllers.Login)
	app.Get("/auth/user/", middlewares.AuthCookie, controllers.GetUserLogin)
	app.Post("/auth/logout/", middlewares.IsLogin, controllers.Logout)
}
