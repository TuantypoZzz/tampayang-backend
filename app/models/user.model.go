package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

// GetUserLoginByEmail retrieves user login information by email with proper error handling
func GetUserLoginByEmail(userEmail string) (entity.UserLogin, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var userLogin entity.UserLogin

	// Input validation
	if userEmail == "" {
		return userLogin, fmt.Errorf("email cannot be empty")
	}

	sqlQuery := "SELECT usr.user_id, usr.user_name, usr.user_email, usr.user_password, usr.user_role FROM users AS usr WHERE usr.user_email = ? LIMIT 1"

	// Log the query attempt (without sensitive data)
	log.Printf("Attempting user lookup for email: %s", maskEmail(userEmail))

	result, err := db.QueryContext(ctx, sqlQuery, userEmail)
	if err != nil {
		log.Printf("Database query error in GetUserLoginByEmail: %v", err)
		return userLogin, fmt.Errorf("database query failed: %w", err)
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(
			&userLogin.User_id,
			&userLogin.User_name,
			&userLogin.User_email,
			&userLogin.User_password,
			&userLogin.User_role)

		if err != nil {
			log.Printf("Row scan error in GetUserLoginByEmail: %v", err)
			return userLogin, fmt.Errorf("failed to scan user data: %w", err)
		}

		log.Printf("User found successfully: %s", userLogin.User_id)
		return userLogin, nil
	}

	// No user found
	log.Printf("No user found for email: %s", maskEmail(userEmail))
	return userLogin, sql.ErrNoRows
}

// maskEmail masks email for logging purposes
func maskEmail(email string) string {
	if len(email) < 3 {
		return "***"
	}
	atIndex := -1
	for i, char := range email {
		if char == '@' {
			atIndex = i
			break
		}
	}
	if atIndex == -1 {
		return "***"
	}

	masked := email[:1] + "***" + email[atIndex:]
	return masked
}
