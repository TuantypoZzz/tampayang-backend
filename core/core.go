package core

import (
	"github.com/gofiber/fiber/v2"
)

// Berisikan koneksion ke db dan handle middleware

func CoreInit(app *fiber.App) {
	
	// INIT MIDLEWARES
	loadMidleWares(app)

}