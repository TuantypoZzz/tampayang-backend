package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func Login(ctx *fiber.Ctx) error {
	// Data from Body POST
	loginReq := new(entity.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Field validation...

	// Validate email dikirim datanya
	if globalFunction.IsEmpty(loginReq.Email) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth003", nil))
	}

	// Validasi password dikirim datanya
	if globalFunction.IsEmpty(loginReq.Password) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth004", nil))
	}

	// validasi format email
	if !globalFunction.IsValidEmail(strings.TrimSpace(loginReq.Email)) {
		response.ErrorResponse(ctx, globalFunction.GetMessage("auth002", nil))
	}

	return nil
}