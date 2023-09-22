package routes

import (
	"github.com/gofiber/fiber/v2"
	example_controller "github.com/nulla-vis/golang-fiber-template/app/controllers/example"
)

func ExampleRoute(app *fiber.App) {
	app.Post("/example/", example_controller.CreateExample)
	app.Get("/example/get_example/:example_id", example_controller.GetExampleById)
	app.Get("/example/all_example/", example_controller.GetAllExample)
	app.Put("/example/update_example/", example_controller.UpdateExample)
}
