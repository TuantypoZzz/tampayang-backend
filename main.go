package main

import (
	"fmt"
	"os"
	"strconv"

	"tampayang-backend/core/middleware"
	_ "tampayang-backend/core/validation"

	"tampayang-backend/app/routes"
	"tampayang-backend/config"
	"tampayang-backend/core"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}

	app := fiber.New()
	// Security middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.SecurityHeaders())
	app.Use(middleware.CORS())
	app.Use(middleware.RateLimiter())
	app.Use(middleware.RequestLogger())
	app.Use(middleware.InputSanitizer())
	app.Use(middleware.BodySizeLimit())
	app.Static("/public", "./public")

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
