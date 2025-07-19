package core

import (
	"fmt"
	"net"
	"os"
	"strconv"

	elasticsearchLib "tampayang-backend/app/libs/elasticsearch"
	"tampayang-backend/config"
	"tampayang-backend/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Berisikan koneksion ke db dan handle middleware

func CoreInit(app *fiber.App) {

	// Initialize middleware
	middlewares.LoadMidleWares(app)

	if isPortInUse(config.PORT) {
		fmt.Printf("\x1b[97;41mPort %s is already in use or invalid.\033[0m \n", config.PORT)
		os.Exit(1)
	}

	// Initialize Elasticsearch connection
	if err := elasticsearchLib.InitElasticsearchClient(); err != nil {
		fmt.Println("\x1b[97;41mElasticsearch FAILED to RUN.\033[0m")
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
