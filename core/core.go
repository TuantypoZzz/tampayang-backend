package core

import (
	"log"

	"github.com/gofiber/fiber/v2"
	elasticsearchLib "github.com/nulla-vis/golang-fiber-template/app/libs/elasticsearch"
)

// Berisikan koneksion ke db dan handle middleware

func CoreInit(app *fiber.App) {
	
	// Initialize middleware
	loadMidleWares(app)

    // Initialize Elasticsearch connection
    if err := elasticsearchLib.InitElasticsearchClient(); err != nil {
        log.Fatalf("Error initializing Elasticsearch: %v", err)
    }
}