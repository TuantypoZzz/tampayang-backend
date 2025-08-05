package helper

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	mylogger "tampayang-backend/core/logger"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret      []byte
	tokenBlacklist = make(map[string]time.Time) // In production, use Redis
)

func init() {
	// Get JWT secret from environment variable or config
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback to config (not recommended for production)
		configData := ConfigJson()
		if dev, ok := configData["development"].(map[string]interface{}); ok {
			if key, ok := dev["key"].(string); ok {
				secret = key
				mylogger.Warn("jwt_secret_from_config", map[string]interface{}{
					"message": "Using JWT secret from config file - not recommended for production",
				})
			}
		}
	}

	if secret == "" {
		mylogger.Error("jwt_secret_missing", map[string]interface{}{
			"message": "JWT secret not found in environment or config",
		})
		panic("JWT secret not configured")
	}

	// Ensure minimum key length for security
	// if len(secret) < 32 {
	// 	mylogger.Error("jwt_secret_too_short", map[string]interface{}{
	// 		"length":  len(secret),
	// 		"minimum": 32,
	// 	})
	// 	panic("JWT secret must be at least 32 characters long")
	// }

	jwtSecret = []byte(secret)
}

// CustomClaims extends jwt.StandardClaims with custom fields
type CustomClaims struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
	JTI       string `json:"jti"` // JWT ID for token revocation
	jwt.StandardClaims
}

// GenerateTokenSecure creates a new JWT token with enhanced security
func GenerateTokenSecure(userID, userName, userEmail, userRole string) (string, error) {
	// Generate unique token ID
	jti, err := generateJTI()
	if err != nil {
		mylogger.Error("jwt_jti_generation_failed", map[string]interface{}{
			"error": err.Error(),
		})
		return "", fmt.Errorf("failed to generate token ID: %w", err)
	}

	// Create claims
	claims := CustomClaims{
		UserID:    userID,
		UserName:  userName,
		UserEmail: userEmail,
		UserRole:  userRole,
		JTI:       jti,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "tampayang-backend",
			Subject:   userID,
			Audience:  "tampayang-frontend",
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // Short expiry
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        jti,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		mylogger.Error("jwt_signing_failed", map[string]interface{}{
			"error":   err.Error(),
			"user_id": userID,
		})
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	mylogger.Info("jwt_token_generated", map[string]interface{}{
		"user_id": userID,
		"jti":     jti,
		"expires": time.Unix(claims.ExpiresAt, 0),
	})

	return tokenString, nil
}

// ValidateTokenSecure validates and parses a JWT token
func ValidateTokenSecure(tokenString string) (*CustomClaims, error) {
	// Check if token is blacklisted
	if isTokenBlacklisted(tokenString) {
		return nil, errors.New("token has been revoked")
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		mylogger.Error("jwt_validation_failed", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Extract claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Additional validation
	if claims.Issuer != "tampayang-backend" {
		return nil, errors.New("invalid token issuer")
	}

	mylogger.Debug("jwt_token_validated", map[string]interface{}{
		"user_id": claims.UserID,
		"jti":     claims.JTI,
	})

	return claims, nil
}

// RevokeToken adds a token to the blacklist
func RevokeToken(tokenString string) error {
	// In production, store this in Redis with TTL
	tokenBlacklist[tokenString] = time.Now()

	mylogger.Info("jwt_token_revoked", map[string]interface{}{
		"token_hash": hashToken(tokenString),
	})

	return nil
}

// isTokenBlacklisted checks if a token is in the blacklist
func isTokenBlacklisted(tokenString string) bool {
	_, exists := tokenBlacklist[tokenString]
	return exists
}

// generateJTI generates a unique JWT ID
func generateJTI() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// hashToken creates a hash of the token for logging (security)
func hashToken(token string) string {
	if len(token) < 10 {
		return "***"
	}
	return token[:8] + "..." + token[len(token)-8:]
}

// Legacy functions for backward compatibility

// GenerateToken creates a JWT token using the old method (deprecated)
func GenerateToken(claims *jwt.MapClaims) string {
	key := fmt.Sprint(string(jwtSecret))
	keyBytes := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString(keyBytes)
	if err != nil {
		panic("helper - GenerateToken, token.SignedString: " + err.Error())
	}
	return webToken
}

// VerfyToken validates a token using the old method (deprecated)
func VerfyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// DecodeToken decodes a token using the old method (deprecated)
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerfyToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("helper - DecodeToken, VerfyToken: %v", err)
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
