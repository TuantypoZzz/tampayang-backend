package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func loadMidleWares(app *fiber.App) {
	// DATABASE CONNECTION
	app.Use(recover.New())
}