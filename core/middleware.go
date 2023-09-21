package core

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func loadMidleWares(app *fiber.App) {

	// Custom middleware to recover from panics and send a custom error response
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			// route := c.Route()
			// fmt.Println(route.Path)
			if r := recover(); r != nil {
				// Recovered from a panic, send a custom error response
				errorMessage := fmt.Sprintf("%v", r) // Create a custom error message
				// Check if r is a map[string]interface{}
				if errMap, ok := r.(map[string]interface{}); ok {
					// Convert the map to JSON
					jsonData, err := json.Marshal(errMap)
					if err != nil {
						// Handle the error if JSON conversion fails
						c.Status(fiber.StatusInternalServerError).JSON(CustomErrorResponse{
							Status:     "error",
							StatusCode: fiber.StatusInternalServerError,
							Payload: map[string]interface{}{
								"error": "Internal Server Error",
							},
						})
						return
					}

					errorMessage = string(jsonData)
				}
				
				c.Status(fiber.StatusInternalServerError).JSON(CustomErrorResponse{
					Status:     "error",
					StatusCode: fiber.StatusInternalServerError,
					Payload: map[string]interface{}{
						"error": errorMessage, // Include the custom error message
					},
				})
			}
		}()
		return c.Next()
	})

}

type CustomErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Payload    interface{} `json:"payload"`
}
