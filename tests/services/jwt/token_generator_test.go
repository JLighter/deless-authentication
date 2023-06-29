package token_generator_test

import (
	"fmt"
	service "glog/services/jwt"
	"os"
	"strconv"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var SECRET = "dummy"
var TOKEN_EXPIRE = 60
var TOKEN_ISSUER = "glog"
var DEFAULT_ID = "1"
var DEFAULT_AUDIENCE = jwt.ClaimStrings{"https://localhost"}

var generator *service.TokenGenerator

func setupValidatorTests() func () {
  generator = service.NewTokenGenerator()
  os.Setenv("SECRET", SECRET)
  os.Setenv("TOKEN_EXPIRE", strconv.Itoa(TOKEN_EXPIRE))
  os.Setenv("TOKEN_ISSUER", TOKEN_ISSUER)

  return func () {
    os.Unsetenv("SECRET")
    os.Unsetenv("TOKEN_EXPIRE")
    os.Unsetenv("TOKEN_ISSUER")
  }
}

func TestGenerateToken(t *testing.T) {

  t.Run("should return non empty token", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    token, _ := generator.GenerateToken(DEFAULT_ID, DEFAULT_AUDIENCE)

    assert.NotEmpty(t, token)
  })

}

func TestTokenClaimSubject(t *testing.T) {

  t.Run("should contains the id passed in first argument", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    expected := "dummy"

    jwtToken, _ := generator.GenerateToken(expected, DEFAULT_AUDIENCE)

    claims := &jwt.RegisteredClaims{}
    jwt.ParseWithClaims(jwtToken, claims, func (token *jwt.Token) (interface{}, error) {
      return []byte(SECRET), nil
    })

    assert.Equal(t, claims.Subject, expected)
  })

}

func TestTokenClaimIssuer(t *testing.T) {

  t.Run("should returns error if TOKEN_ISSUER environment variable is not set", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    os.Unsetenv("TOKEN_ISSUER")

    _, err := generator.GenerateToken(DEFAULT_ID, DEFAULT_AUDIENCE)

    expectedError := fmt.Errorf("cannot generate claim: TOKEN_ISSUER environment variable is not set")

    if assert.Error(t, err) {
      assert.Equal(t, expectedError, err)
    }
  })

  t.Run("should set issuer from TOKEN_ISSUER environment variable", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    expected := TOKEN_ISSUER

    jwtToken, _ := generator.GenerateToken(expected, DEFAULT_AUDIENCE)

    claims := &jwt.RegisteredClaims{}
    jwt.ParseWithClaims(jwtToken, claims, func (token *jwt.Token) (interface{}, error) {
      return []byte(SECRET), nil
    })

    assert.Equal(t, claims.Issuer, expected)
  })

}


func TestTokenCypher(t *testing.T) {

  t.Run("should returns error if SECRET environment variable is not set", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    os.Unsetenv("SECRET")

    _, err := generator.GenerateToken(DEFAULT_ID, DEFAULT_AUDIENCE)

    expectedError := fmt.Errorf("cannot generate token: SECRET environment variable is not set")

    if assert.Error(t, err) {
      assert.Equal(t, expectedError, err)
    }
  })

  t.Run("should return a valid token signed using environment variable 'SECRET'", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    jwtToken, _ := generator.GenerateToken(DEFAULT_ID, DEFAULT_AUDIENCE)

    claims := &jwt.RegisteredClaims{}

    token, _ := jwt.ParseWithClaims(jwtToken, claims, func (token *jwt.Token) (interface{}, error) {
      return []byte(SECRET), nil
    })

    assert.True(t, token.Valid)
  })

  t.Run("should cypher token with HS256 method", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    jwtToken, _ := generator.GenerateToken(DEFAULT_ID, DEFAULT_AUDIENCE)

    claims := &jwt.RegisteredClaims{}

    token, _ := jwt.ParseWithClaims(jwtToken, claims, func (token *jwt.Token) (interface{}, error) {
      return []byte(SECRET), nil
    })

    assert.Equal(t, token.Method.Alg(), "HS256") 
  })
}

func TestTokenAudience(t *testing.T) {

  t.Run("should set audience from second passed parameter", func(t *testing.T) {
    teardown := setupValidatorTests()
    defer teardown()

    expected := jwt.ClaimStrings{"https://localhost"}

    jwtToken, _ := generator.GenerateToken(DEFAULT_ID, expected)

    claims := &jwt.RegisteredClaims{}
    jwt.ParseWithClaims(jwtToken, claims, func (token *jwt.Token) (interface{}, error) {
      return []byte(SECRET), nil
    })

    assert.Equal(t, claims.Audience, expected)
  })
}
