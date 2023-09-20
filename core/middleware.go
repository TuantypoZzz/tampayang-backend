package core

import (
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
