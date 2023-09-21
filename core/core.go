package core

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	elasticsearchLib "github.com/nulla-vis/golang-fiber-template/app/libs/elasticsearch"
	"github.com/nulla-vis/golang-fiber-template/config"
)

// Berisikan koneksion ke db dan handle middleware

func CoreInit(app *fiber.App) {

	if isPortInUse(config.PORT) {
        fmt.Printf("\x1b[97;41mPort %s is already in use or invalid.\033[0m \n", config.PORT)
        os.Exit(1)
    }
	
	// Initialize middleware
	loadMidleWares(app)

    // Initialize Elasticsearch connection
    if err := elasticsearchLib.InitElasticsearchClient(); err != nil {
        log.Fatalf("Error initializing Elasticsearch: %v", err)
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