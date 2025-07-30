package controllers

import (
	"strings"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/config/constant"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/helper"
	"tampayang-backend/core/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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
		Secure:   true,
		SameSite: "None",
		Expires:  time.Now().Add(time.Minute * constant.TOKEN_EXPIRE_MINUTE),
	})

	result := entity.ResultToken{
		Access_token: token,
	}

	return response.SuccessResponse(ctx, result)
}

func GetUserLogin(ctx *fiber.Ctx) error {
	// access userInfo object from jwt.MapClaims------------------
	userInfo := ctx.Locals("userInfo").(jwt.MapClaims)

	if globalFunction.IsEmpty(userInfo) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth001", nil))
	}
	userLogin := entity.LoggedInUser{
		User_id:    userInfo["user_id"].(string),
		User_name:  userInfo["user_name"].(string),
		User_email: userInfo["user_email"].(string),
		User_role:  userInfo["user_role"].(string),
	}

	return response.SuccessResponse(ctx, userLogin)
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
		Secure:   true,
		SameSite: "None",
	})

	return response.SuccessResponse(ctx, globalFunction.GetMessage("success", nil))
}
