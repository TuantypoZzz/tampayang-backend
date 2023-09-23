package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
	"github.com/nulla-vis/golang-fiber-template/core/middlewares"
)

func ExampleRoute(app *fiber.App) {
	app.Post("/example/", middlewares.AuthCookie, controllers.CreateExample)
	app.Get("/example/all_example/", middlewares.AuthCookie, controllers.GetAllExample)
	app.Get("/example/get_example/:example_id", middlewares.AuthCookie, controllers.GetExampleById)
	app.Put("/example/update_example/", middlewares.AuthCookie, controllers.UpdateExample)
	app.Delete("/example/delete_example/:example_id", middlewares.AuthCookie, controllers.DeleteExample)
}