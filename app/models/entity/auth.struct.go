package entity

import (
	"regexp"
	"strings"
	"unicode"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Validate validates the login request
func (lr *LoginRequest) Validate() error {
	// Trim whitespace
	lr.Email = strings.TrimSpace(lr.Email)
	lr.Password = strings.TrimSpace(lr.Password)

	// Email validation
	if lr.Email == "" {
		return NewValidationError("email", "Email is required")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, lr.Email); !matched {
		return NewValidationError("email", "Invalid email format")
	}

	// Password validation
	if lr.Password == "" {
		return NewValidationError("password", "Password is required")
	}

	if len(lr.Password) < 8 {
		return NewValidationError("password", "Password must be at least 8 characters long")
	}

	// Check for password complexity
	// if !isStrongPassword(lr.Password) {
	// 	return NewValidationError("password", "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	// }

	return nil
}

// UserLogin represents user data from database (internal use only)
type UserLogin struct {
	User_id       string `json:"-"` // Never expose in JSON
	User_name     string `json:"-"`
	User_email    string `json:"-"`
	User_password string `json:"-"` // NEVER expose password hash
	User_role     string `json:"-"`
}

// ToLoggedInUser converts UserLogin to safe LoggedInUser for API responses
func (ul *UserLogin) ToLoggedInUser() LoggedInUser {
	return LoggedInUser{
		User_id:    ul.User_id,
		User_name:  ul.User_name,
		User_email: ul.User_email,
		User_role:  ul.User_role,
	}
}

type LoggedInUser struct {
	User_id    string `json:"user_id"`
	User_name  string `json:"user_name"`
	User_email string `json:"user_email"`
	User_role  string `json:"user_role"`
}

type ResultToken struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int64  `json:"expires_in"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// isStrongPassword checks if password meets complexity requirements
func isStrongPassword(password string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
