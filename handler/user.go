package handler

import (
	"context"
	"glog/services/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var ctx = context.Background()

func GetUser(c *fiber.Ctx) error {
  token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

  db := database.NewDatabase()
  user, err := db.GetUser(email)

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

  db := database.NewDatabase()
  err := db.RegisterUser(database.NewUser{
    Email: c.FormValue("email"),
    Username: c.FormValue("username"),
    Password: c.FormValue("password"),
    Admin: false,
  })

  if err != nil {
    log.Printf("Error registering user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.SendStatus(fiber.StatusOK)
}

func UpdateUser(c *fiber.Ctx) error {

  newUser := database.User{
    Email: c.FormValue("email"),
    Username: c.FormValue("username"),
  }

  db := database.NewDatabase()
  err := db.UpdateUser(newUser)

  if err != nil {
    log.Printf("Error updating user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.Status(fiber.StatusOK).JSON(newUser)
}

func ChangePassword(c *fiber.Ctx) error {
  email := ""
  password := c.FormValue("password")

  db := database.NewDatabase()
  err := db.ChangePassword(email, password)

  if err != nil {
    log.Printf("Error changing password: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}
