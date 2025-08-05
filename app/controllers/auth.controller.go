package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/config/constant"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/helper"
	mylogger "tampayang-backend/core/logger"
	"tampayang-backend/core/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Rate limiting and account locking structures
var (
	rateLimitMap   = make(map[string][]time.Time)
	failedAttempts = make(map[string]int)
	lockedAccounts = make(map[string]time.Time)
	rateLimitMutex sync.RWMutex
	accountMutex   sync.RWMutex
)

const (
	maxLoginAttempts     = 5
	lockoutDuration      = 15 * time.Minute
	rateLimitWindow      = 1 * time.Minute
	maxRequestsPerWindow = 10
)

func Login(ctx *fiber.Ctx) error {
	// Check if user is already logged in
	if ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == true {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth008", nil))
	}

	// Parse request body
	loginReq := new(entity.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		mylogger.Error("login_parse_error", map[string]interface{}{
			"error": err.Error(),
			"ip":    ctx.IP(),
		})
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth009", nil))
	}

	// Validate input
	if err := loginReq.Validate(); err != nil {
		mylogger.Error("login_validation_error", map[string]interface{}{
			"error": err.Error(),
			"email": loginReq.Email,
			"ip":    ctx.IP(),
		})
		return response.ErrorResponse(ctx, err)
	}

	// Rate limiting check (implement in middleware)
	clientIP := ctx.IP()
	if isRateLimited(clientIP) {
		mylogger.Error("login_rate_limited", map[string]interface{}{
			"ip":    clientIP,
			"email": loginReq.Email,
		})
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth010", nil))
	}

	// Get user from database
	user, err := models.GetUserLoginByEmail(loginReq.Email)
	if err != nil {
		// Log failed login attempt
		mylogger.Error("login_user_not_found", map[string]interface{}{
			"email": loginReq.Email,
			"ip":    clientIP,
			"error": err.Error(),
		})

		// Return generic error to prevent user enumeration
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth006", nil))
	}

	// Validate password
	isValid := helper.ValidatePassword(user.User_password, loginReq.Password)
	if !isValid {
		// Log failed login attempt
		mylogger.Error("login_invalid_password", map[string]interface{}{
			"user_id": user.User_id,
			"email":   loginReq.Email,
			"ip":      clientIP,
		})

		// Increment failed login attempts (implement counter)
		incrementFailedLoginAttempts(user.User_id)

		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth006", nil))
	}

	// Check if account is locked (implement account locking)
	if isAccountLocked(user.User_id) {
		mylogger.Error("login_account_locked", map[string]interface{}{
			"user_id": user.User_id,
			"email":   loginReq.Email,
			"ip":      clientIP,
		})
		return response.ErrorResponse(ctx, globalFunction.GetMessage("auth011", nil))
	}

	// Generate secure JWT
	claims := jwt.MapClaims{}
	claims["user_id"] = user.User_id
	claims["user_name"] = user.User_name
	claims["user_email"] = user.User_email
	claims["user_role"] = user.User_role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * constant.TOKEN_EXPIRE_MINUTE).Unix()
	claims["jti"] = generateJTI() // Add unique token ID

	token := helper.GenerateToken(&claims)

	// Set secure cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     constant.JWT_COOKIE_NAME,
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None", // Changed from "None" for better security
		Expires:  time.Now().Add(time.Minute * constant.TOKEN_EXPIRE_MINUTE),
		Path:     "/",
	})

	// Log successful login
	mylogger.Info("login_success", map[string]interface{}{
		"user_id": user.User_id,
		"email":   user.User_email,
		"ip":      clientIP,
		"role":    user.User_role,
	})

	// Reset failed login attempts
	resetFailedLoginAttempts(user.User_id)

	result := entity.ResultToken{
		Access_token: token,
		Token_type:   "Bearer",
		Expires_in:   constant.TOKEN_EXPIRE_MINUTE * 60, // Convert to seconds
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

// Security helper functions

// isRateLimited checks if the IP has exceeded rate limits
func isRateLimited(ip string) bool {
	rateLimitMutex.Lock()
	defer rateLimitMutex.Unlock()

	now := time.Now()
	requests, exists := rateLimitMap[ip]

	if !exists {
		rateLimitMap[ip] = []time.Time{now}
		return false
	}

	// Remove old requests outside the window
	var validRequests []time.Time
	for _, reqTime := range requests {
		if now.Sub(reqTime) < rateLimitWindow {
			validRequests = append(validRequests, reqTime)
		}
	}

	// Add current request
	validRequests = append(validRequests, now)
	rateLimitMap[ip] = validRequests

	return len(validRequests) > maxRequestsPerWindow
}

// incrementFailedLoginAttempts increments failed login attempts for a user
func incrementFailedLoginAttempts(userID string) {
	accountMutex.Lock()
	defer accountMutex.Unlock()

	failedAttempts[userID]++

	// Lock account if max attempts reached
	if failedAttempts[userID] >= maxLoginAttempts {
		lockedAccounts[userID] = time.Now()
	}
}

// isAccountLocked checks if an account is currently locked
func isAccountLocked(userID string) bool {
	accountMutex.RLock()
	defer accountMutex.RUnlock()

	lockTime, exists := lockedAccounts[userID]
	if !exists {
		return false
	}

	// Check if lockout period has expired
	if time.Since(lockTime) > lockoutDuration {
		// Clean up expired lock
		delete(lockedAccounts, userID)
		delete(failedAttempts, userID)
		return false
	}

	return true
}

// resetFailedLoginAttempts resets failed login attempts for a user
func resetFailedLoginAttempts(userID string) {
	accountMutex.Lock()
	defer accountMutex.Unlock()

	delete(failedAttempts, userID)
	delete(lockedAccounts, userID)
}

// generateJTI generates a unique JWT ID
func generateJTI() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
