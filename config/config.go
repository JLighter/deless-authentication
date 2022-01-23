package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
    return "", fmt.Errorf("error loading .env file : %s", err)
	}
	return os.Getenv(key), nil
}
