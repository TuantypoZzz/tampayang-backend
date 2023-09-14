package main

import (
	"github.com/nulla-vis/golang-fiber-template/app/routes"
	"github.com/nulla-vis/golang-fiber-template/core"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app:= fiber.New()

	// INITIALIZE CORE
	core.CoreInit(app)

	// INITIALIZE ROUTE
	routes.RouteInit(app)

	app.Listen((":3000"))
}