package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/controllers"
)

func ExampleRoute(app *fiber.App) {
	app.Post("/example/", controllers.CreateExample)
	app.Get("/example/get_example/:example_id", controllers.GetExampleById)
	app.Get("/example/all_example/", controllers.GetAllExample)
	app.Put("/example/update_example/", controllers.UpdateExample)
}


func Middleware(ctx *fiber.Ctx) error {

	// headers authorization
	token := ctx.Get("x-token")
	if token == "" || token != "secret" {
		panic("Authorization Failed")
	}
	
	return ctx.Next()
}
