package middleware

import (
	"strings"
	"sync"
	"time"

	"tampayang-backend/config"
	mylogger "tampayang-backend/core/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Rate limiting storage
var (
	rateLimitMap   = make(map[string][]time.Time)
	rateLimitMutex sync.RWMutex
)

// SecurityHeaders adds security headers to all responses
func SecurityHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Set security headers manually since helmet might not be available
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; child-src 'none'; worker-src 'none'; frame-ancestors 'none'; form-action 'self'; base-uri 'self';")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		return c.Next()
	}
}

// CORS configures Cross-Origin Resource Sharing
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     config.ALLOWED_ORIGINS, // Use configurable origins from environment
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	})
}

// RateLimiter implements custom rate limiting
func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()

		rateLimitMutex.Lock()
		defer rateLimitMutex.Unlock()

		now := time.Now()
		requests, exists := rateLimitMap[ip]

		if !exists {
			rateLimitMap[ip] = []time.Time{now}
			return c.Next()
		}

		// Remove old requests outside the window (1 minute)
		var validRequests []time.Time
		for _, reqTime := range requests {
			if now.Sub(reqTime) < 1*time.Minute {
				validRequests = append(validRequests, reqTime)
			}
		}

		// Check if limit exceeded (100 requests per minute)
		if len(validRequests) >= 100 {
			mylogger.Security("rate_limit_exceeded", map[string]interface{}{
				"ip":         ip,
				"user_agent": c.Get("User-Agent"),
				"path":       c.Path(),
				"method":     c.Method(),
			})
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Rate limit exceeded",
				"message": "Too many requests, please try again later",
			})
		}

		// Add current request
		validRequests = append(validRequests, now)
		rateLimitMap[ip] = validRequests

		return c.Next()
	}
}

// AuthRateLimiter implements stricter rate limiting for auth endpoints
func AuthRateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		key := "auth_" + ip

		rateLimitMutex.Lock()
		defer rateLimitMutex.Unlock()

		now := time.Now()
		requests, exists := rateLimitMap[key]

		if !exists {
			rateLimitMap[key] = []time.Time{now}
			return c.Next()
		}

		// Remove old requests outside the window (15 minutes)
		var validRequests []time.Time
		for _, reqTime := range requests {
			if now.Sub(reqTime) < 15*time.Minute {
				validRequests = append(validRequests, reqTime)
			}
		}

		// Check if limit exceeded (5 requests per 15 minutes)
		if len(validRequests) >= 5 {
			mylogger.Security("auth_rate_limit_exceeded", map[string]interface{}{
				"ip":         ip,
				"user_agent": c.Get("User-Agent"),
				"path":       c.Path(),
				"method":     c.Method(),
			})
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Authentication rate limit exceeded",
				"message": "Too many login attempts, please try again in 15 minutes",
			})
		}

		// Add current request
		validRequests = append(validRequests, now)
		rateLimitMap[key] = validRequests

		return c.Next()
	}
}

// RequestID adds unique request ID to each request
func RequestID() fiber.Handler {
	return requestid.New()
}

// RequestLogger logs all incoming requests
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Log request details
		duration := time.Since(start)

		logData := map[string]interface{}{
			"method":      c.Method(),
			"path":        c.Path(),
			"status_code": c.Response().StatusCode(),
			"duration_ms": duration.Milliseconds(),
			"ip":          c.IP(),
			"user_agent":  c.Get("User-Agent"),
			"request_id":  c.Locals("requestid"),
		}

		// Log based on status code
		if c.Response().StatusCode() >= 400 {
			mylogger.Error("http_request_error", logData)
		} else {
			mylogger.Info("http_request", logData)
		}

		// Log performance metrics for slow requests
		if duration > 1*time.Second {
			mylogger.Performance("slow_request", duration, logData)
		}

		return err
	}
}

// InputSanitizer sanitizes input data
func InputSanitizer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Sanitize headers
		userAgent := c.Get("User-Agent")
		if containsSuspiciousPatterns(userAgent) {
			mylogger.Security("suspicious_user_agent", map[string]interface{}{
				"ip":         c.IP(),
				"user_agent": userAgent,
				"path":       c.Path(),
			})
		}

		// Check for suspicious query parameters
		c.Context().QueryArgs().VisitAll(func(key, value []byte) {
			if containsSuspiciousPatterns(string(value)) {
				mylogger.Security("suspicious_query_param", map[string]interface{}{
					"ip":    c.IP(),
					"key":   string(key),
					"value": string(value),
					"path":  c.Path(),
				})
			}
		})

		return c.Next()
	}
}

// containsSuspiciousPatterns checks for common attack patterns
func containsSuspiciousPatterns(input string) bool {
	suspiciousPatterns := []string{
		"<script",
		"javascript:",
		"onload=",
		"onerror=",
		"eval(",
		"alert(",
		"document.cookie",
		"union select",
		"drop table",
		"insert into",
		"delete from",
		"update set",
		"../",
		"..\\",
		"cmd.exe",
		"/bin/sh",
		"passwd",
		"/etc/",
	}

	lowerInput := strings.ToLower(input)
	for _, pattern := range suspiciousPatterns {
		if strings.Contains(lowerInput, pattern) {
			return true
		}
	}

	return false
}

// BodySizeLimit limits request body size using Fiber's built-in method
func BodySizeLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check body size (10MB limit)
		if len(c.Body()) > 10*1024*1024 {
			return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
				"error":   "Request body too large",
				"message": "Request body must be less than 10MB",
			})
		}
		return c.Next()
	}
}
