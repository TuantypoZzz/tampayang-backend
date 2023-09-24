package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nulla-vis/golang-fiber-template/app/routes"
	"github.com/nulla-vis/golang-fiber-template/config"
	"github.com/nulla-vis/golang-fiber-template/core"
	"github.com/nulla-vis/golang-fiber-template/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}

	app:= fiber.New()

	// INITIALIZE CORE
	core.CoreInit(app)

	// KEEP TRACK OF REGISTERED ROUTES
	registeredRoutes := make(map[string]bool)
	// INITIALIZE ROUTE
	routes.RouteInit(app)
	// ROUTE VALIDATION (ONLY DEV ENV)
	if config.GO_ENV == "development" {
		middlewares.RouteValidation(app, registeredRoutes)
	}

    port, err := strconv.Atoi(config.PORT)
    if err != nil {
        fmt.Printf("\x1b[97;41mInvalid port number: %v\033[0m\n", err)
        os.Exit(1)
    }

	// Start the Fiber server
	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
        fmt.Printf("Error starting the server: %v\n", err)
        os.Exit(1)
    }
}