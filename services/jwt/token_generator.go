package jwt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenGenerator struct { }

func NewTokenGenerator() *TokenGenerator {
  return &TokenGenerator{}
}

func (t *TokenGenerator) createClaim(id string, audiance []string) (*jwt.RegisteredClaims, error) {
	token_expire := os.Getenv("TOKEN_EXPIRE")

	expire, err := strconv.Atoi(token_expire)
	if err != nil {
    log.Printf("cannot convert TOKEN_EXPIRE to integer : %s", err)
	}

	token_issuer := os.Getenv("TOKEN_ISSUER")
  if token_issuer == "" {
    return nil, fmt.Errorf("TOKEN_ISSUER environment variable is not set")
  }

	return &jwt.RegisteredClaims{
    Issuer:    token_issuer,
		Subject:   id,
		Audience:  audiance,
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expire))),
		NotBefore: jwt.NewNumericDate(time.Now()),
    IssuedAt:  jwt.NewNumericDate(time.Now()),
    ID:        uuid.NewString(),
	}, nil
}

func (t *TokenGenerator) GenerateToken(id string, audiance []string) (string, error) {
	claim, err := t.createClaim(id, audiance)
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

