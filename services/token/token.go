package token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
  "github.com/google/uuid"
)

var DEFAULT_EXPIRATION_TIME = 3600

func createClaim(id string) (*jwt.RegisteredClaims, error) {
	token_expire := os.Getenv("TOKEN_EXPIRE")

	expire, err := strconv.Atoi(token_expire)
	if err != nil {
    log.Printf("cannot convert TOKEN_EXPIRE to integer : %s (fallback to DEFAULT:3600)", err)
    expire = DEFAULT_EXPIRATION_TIME
	}

	return &jwt.RegisteredClaims{
		Issuer:    "glog-authentication",
		Subject:   id,
		Audience:  []string{""},
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expire))),
		NotBefore: jwt.NewNumericDate(time.Now()),
    IssuedAt:  jwt.NewNumericDate(time.Now()),
    ID:        uuid.NewString(),
	}, nil
}

func GenerateToken(id string, admin bool) (string, error) {
	claim, err := createClaim(id)
	if err != nil {
		return "", fmt.Errorf("cannot generate claim: %s", err)
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secret := os.Getenv("SECRET")
  if secret == "" {
    return "", fmt.Errorf("cannot generate token: SECRET environment variable is not set")
  }

	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("cannot sign token: %s", err)
	}

	return tokenString, nil
}
