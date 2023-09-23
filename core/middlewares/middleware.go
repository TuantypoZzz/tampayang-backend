package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/config"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func LoadMidleWares(app *fiber.App) {

	// Custom middleware to recover from panics and send a custom error response
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
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

			if config.GO_ENV == "development" {
				fmt.Print("Request Route:", c.Path())
				switch statusCode := c.Response().StatusCode(); statusCode {
				case fiber.StatusOK:
					fmt.Println(" - \033[32m200\033[0m")
				case fiber.StatusInternalServerError:
					fmt.Println(" - \033[31m500\033[0m")
				case fiber.StatusNotFound:
					fmt.Println(" - \033[33m404\033[0m")
				}
			}
		}()
		return c.Next()
	})
}

func RouteValidation(app *fiber.App, registeredRoutes map[string]bool) {
	// Custom middleware to check if the current path is a valid route
	app.Use(func(c *fiber.Ctx) error {
		// Check if the current path is a valid route
		if _, exists := registeredRoutes[c.Path()]; !exists {
			// Handle the case where the path is not a valid route
			fmt.Print("(ROUTE NOT FOUND) ")
			return response.ErrorResponse(c, globalFunction.GetMessage("err003", nil))
		}
		// Handle the case where the path is a valid route
		return c.Next()
	})
}

func Auth(ctx *fiber.Ctx) error {
	// headers authorization
	token := ctx.Get("x-token")
	if token == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}

	// _, err := helper.VerfyToken(token)
	claims, err := helper.DecodeToken(token)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}

	role := claims["user_role"].(string)
	if role != constant.ADMIN_ROLE {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth007", nil))
	}

	// Unix timestamp in scientific notation
	scientificNotation := claims["expire"].(float64)

	// Convert scientific notation to a regular Unix timestamp
	unixTimestamp := int64(scientificNotation)

	// Get the current Unix timestamp
	currentTime := time.Now().Unix()

	// Check if the timestamp has passed the current time
	if unixTimestamp < currentTime {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}

	ctx.Locals("userInfo", claims)
	ctx.Locals("user_role", claims["user_role"])
	
	
	return ctx.Next()
}





type CustomErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Payload    interface{} `json:"payload"`
}
