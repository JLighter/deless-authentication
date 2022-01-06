package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) (string, error) {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
    return "", fmt.Errorf("error loading .env file : %s", err)
	}
	return os.Getenv(key), nil
}
