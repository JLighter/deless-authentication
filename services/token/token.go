package token

import (
	"glog/config"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createClaim(email string, admin bool) (*jwt.MapClaims, error) {
  token_expire, err := config.Config("TOKEN_EXPIRE")
  if err != nil {
    log.Fatalf("cannot get TOKEN_EXPIRE variable : %s", err)
  }

  expire, err := strconv.Atoi(token_expire)
  if err != nil {
    log.Fatalf("cannot convert TOKEN_EXPIRE to integer : %s", err)
  }

  return &jwt.MapClaims{
		"email":  email,
		"admin": admin,
		"exp":   time.Now().Add(time.Second * time.Duration(expire)).Unix(),
	}, nil
}

func GenerateToken(email string, admin bool) string {
  claim, err := createClaim(email, admin)
  if err != nil {
    log.Fatalf("cannot generate claim: %s", err)
  }

  newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

  secret, err := config.Config("SECRET")
	if err != nil {
    log.Fatalf("cannot get secret: %s", err)
	}

	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
    log.Fatalf("cannot sign token: %s", err)
	}

  return tokenString
}
