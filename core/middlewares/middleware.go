package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"tampayang-backend/config/constant"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/helper"
	mylogger "tampayang-backend/core/logger"
	"tampayang-backend/core/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func LoadMidleWares(app *fiber.App) {
	// Cors Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5173",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

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

				// LOGGER HERE -----
				logData := response.GetLogData(c, errorMessage)
				mylogger.Error("panic_error", logData)
				// LOGGER END ------

				c.Status(fiber.StatusInternalServerError).JSON(CustomErrorResponse{
					Status:     "error",
					StatusCode: fiber.StatusInternalServerError,
					Payload: map[string]interface{}{
						"error": errorMessage,
					},
				})
			}

			if os.Getenv("GO_ENV") == "development" {
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

	if ctx.Locals("userInfo") == nil {
		ctx.Locals("userInfo", claims)
	}

	return ctx.Next()
}

func AuthCookie(ctx *fiber.Ctx) error {
	token := ctx.Cookies(constant.JWT_COOKIE_NAME)

	if token == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}

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

	return ctx.Next()
}

func IsLogin(ctx *fiber.Ctx) error {
	// Set initial value for isLogin validation
	isLogin := false
	// Check if user already has a valid JWT cookie
	jwtCookie := ctx.Cookies(constant.JWT_COOKIE_NAME)
	if jwtCookie != "" {
		// User is already logged in
		isLogin = true
	}

	// Parse the expiration time from the cookie
	expirationTime, _ := time.Parse(time.RFC3339, jwtCookie)

	// Check if the cookie has expired
	if time.Now().Before(expirationTime) {
		// Cookie not yet expired
		isLogin = true
	}

	ctx.Locals("isLogin", isLogin)

	// Continue processing the request
	return ctx.Next()
}

type CustomErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Payload    interface{} `json:"payload"`
}
