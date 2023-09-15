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
    GO_ENV		string
	PORT 		string
)

func LoadEnvVariables() error {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        return err
    }

    // Assign environment variables to variables
    GO_ENV = os.Getenv("GO_ENV")
    PORT   = os.Getenv("PORT")


    return nil
}