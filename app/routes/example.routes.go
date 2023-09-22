package routes

import (
	"github.com/gofiber/fiber/v2"
	example_controller "github.com/nulla-vis/golang-fiber-template/app/controllers/example"
)

func ExampleRoute(app *fiber.App) {
	app.Post("/example/", example_controller.CreateExample)
}