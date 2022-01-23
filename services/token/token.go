package token

import (
	"fmt"
	"glog/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createClaim(id string) (*jwt.MapClaims, error) {
	token_expire, err := config.Config("TOKEN_EXPIRE")
	if err != nil {
		return nil, fmt.Errorf("cannot get TOKEN_EXPIRE variable : %s", err)
	}

	expire, err := strconv.Atoi(token_expire)
	if err != nil {
		return nil, fmt.Errorf("cannot convert TOKEN_EXPIRE to integer : %s", err)
	}

	return &jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(expire)).Unix(),
	}, nil
}

func GenerateToken(id string, admin bool) (string, error) {
	claim, err := createClaim(id)
	if err != nil {
		return "", fmt.Errorf("cannot generate claim: %s", err)
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secret, err := config.Config("SECRET")
	if err != nil {
		return "", fmt.Errorf("cannot get secret: %s", err)
	}

	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("cannot sign token: %s", err)
	}

	return tokenString, nil
}
