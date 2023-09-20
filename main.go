package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/nulla-vis/golang-fiber-template/app/routes"
	"github.com/nulla-vis/golang-fiber-template/config"
	"github.com/nulla-vis/golang-fiber-template/core"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}

	app:= fiber.New()

	// INITIALIZE CORE
	core.CoreInit(app)

	// INITIALIZE ROUTE
	routes.RouteInit(app)
	
	if isPortInUse(config.PORT) {
        fmt.Printf("\x1b[97;41mPort %s is already in use or invalid.\033[0m \n", config.PORT)
        os.Exit(1)
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

func isPortInUse(port string) bool {
    portInt, err := strconv.Atoi(port)
    if err != nil {
        fmt.Printf("Invalid port number: %v\n", err)
        return true
    }

    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", portInt))
    if err != nil {
        return true
    }
    defer listener.Close()
    return false
}