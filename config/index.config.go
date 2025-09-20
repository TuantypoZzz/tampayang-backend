package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	// GET CURRENT FILE FULL PATH FROM RUNTIME
	_, b, _, _ = runtime.Caller(0)

	// ROOT FOLDER OF THIS PROJECT
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

var (
	// Declare variables for environment variables
	GO_ENV string
	PORT   string

	// CORS configuration
	ALLOWED_ORIGINS string

	// Deklarasikan variabel baru untuk WhatsApp
	WHATSAPP_API_URL       string
	WHATSAPP_API_KEY       string
	WHATSAPP_SENDER_NUMBER string

	EMAIL_SENDER_NAME     string
	EMAIL_SENDER_ADDRESS  string
	EMAIL_SENDER_PASSWORD string
	SMTP_HOST             string
	SMTP_PORT             string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
)

func LoadEnvVariables() error {
	// Load environment variables from the .env file
	if err := godotenv.Load(filepath.Join(ProjectRootPath, ".env")); err != nil {
		return err
	}

	// Assign environment variables to variables
	GO_ENV = os.Getenv("GO_ENV")
	PORT = os.Getenv("PORT")

	// CORS configuration - fallback to development defaults if not set
	ALLOWED_ORIGINS = os.Getenv("ALLOWED_ORIGINS")
	if ALLOWED_ORIGINS == "" {
		if GO_ENV == "development" {
			ALLOWED_ORIGINS = "http://localhost:5173,http://127.0.0.1:5173,http://localhost:3000,http://127.0.0.1:3000"
		} else {
			ALLOWED_ORIGINS = "https://yourdomain.com" // Update with your production domain
		}
	}

	// Ambil variabel baru dari .env
	WHATSAPP_API_URL = os.Getenv("WHATSAPP_API_URL")
	WHATSAPP_API_KEY = os.Getenv("WHATSAPP_API_KEY")
	WHATSAPP_SENDER_NUMBER = os.Getenv("WHATSAPP_SENDER_NUMBER")

	EMAIL_SENDER_NAME = os.Getenv("EMAIL_SENDER_NAME")
	EMAIL_SENDER_ADDRESS = os.Getenv("EMAIL_SENDER_ADDRESS")
	EMAIL_SENDER_PASSWORD = os.Getenv("EMAIL_SENDER_PASSWORD")
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = os.Getenv("SMTP_PORT")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	return nil
}
