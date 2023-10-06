package controllers

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/models"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func Login(ctx *fiber.Ctx) error {
	if ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == true {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth008", nil))
	}

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
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth002", nil))
	}

	// Get user from db
	user := models.GetUserLoginByEmail(loginReq.Email)
	// validasi data user terdaftar (menggunakan id)
	if globalFunction.IsEmpty(user.User_id) {
		replacements := []interface{}{loginReq.Email}
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth005", replacements))
	}

	// validasi password
	isValid := helper.ValidatePassword(user.User_password, loginReq.Password)
	if !isValid {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth006", nil))
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	// Insert user data to claims
	claims["user_id"] = user.User_id
	claims["user_name"] = user.User_name
	claims["user_email"] = user.User_email
	claims["user_role"] = user.User_role
	claims["expire"] = time.Now().Add(time.Minute * constant.TOKEN_EXPIRE_MINUTE).Unix()

	token := helper.GenerateToken(&claims)

	// Set the JWT token as a cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     constant.JWT_COOKIE_NAME,
		Value:    token,
		HTTPOnly: true,
		SameSite: "Strict",
		Expires:  time.Now().Add(time.Minute * constant.TOKEN_EXPIRE_MINUTE),
	})

	result := entity.ResultToken{
		Access_token: token,
	}

	return response.SuccessResponse(ctx, result)
}

func Logout(ctx *fiber.Ctx) error {
	if ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == false {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}

	// Remove the JWT cookie to log the user out
	ctx.Cookie(&fiber.Cookie{
		Name:     constant.JWT_COOKIE_NAME,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expire the cookie immediately
		HTTPOnly: true,
		SameSite: "Strict",
	})

	return response.SuccessResponse(ctx, globalFunction.GetMessage("success", nil))
}
