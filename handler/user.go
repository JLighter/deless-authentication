package handler

import (
	"context"
	"fmt"
	"glog/services/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var ctx = context.Background()

func getUserId(c *fiber.Ctx) (string, error) {
  token := c.Locals("user").(*jwt.Token)
  claim := token.Claims.(jwt.MapClaims)
  if claim["id"] == nil {
    return "", fmt.Errorf("invalid token")
  } else {
    return claim["id"].(string), nil
  }
}

func GetUser(c *fiber.Ctx) error {
  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  db := database.GetMongoDB()
  user, err := db.GetUserById(userId)

  if err != nil {
    log.Printf("Error getting user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if user == nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  return c.Status(fiber.StatusOK).JSON(user)
}

func RegisterUser(c *fiber.Ctx) error {

  db := database.GetMongoDB()

  if exists := db.UserExists(c.FormValue("email")); exists == true {
    return c.SendStatus(fiber.StatusOK)
  }

  id := uuid.New().String()
  err := db.RegisterUser(database.User{
    Id: id,
    Email: c.FormValue("email"),
    Username: c.FormValue("username"),
    Admin: false,
  })

  db.SetPassword(database.Password{
    UserId: id,
    Value: c.FormValue("password"),
  })

  if err != nil {
    log.Printf("Error registering user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.SendStatus(fiber.StatusOK)
}

func UpdateSelf(c *fiber.Ctx) error {

  db := database.GetMongoDB()
  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  user, err := db.GetUserById(userId)
  if err != nil {
    log.Printf("Error getting user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if user == nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  newUser := database.User{
    Id: userId,
    Email: c.FormValue("email"),
    Username: c.FormValue("username"),
    Admin: user.Admin,
  }

  err = db.UpdateUser(newUser)
  if err != nil {
    log.Printf("Error updating user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.Status(fiber.StatusOK).JSON(newUser)
}

func ChangePassword(c *fiber.Ctx) error {
  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }
  password := c.FormValue("password")

  db := database.GetMongoDB()
  err = db.ChangePassword(database.Password{
    UserId: userId,
    Value: password,
  })

  if err != nil {
    log.Printf("Error changing password: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}
