package handler

import (
	"context"
	"glog/services/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var ctx = context.Background()

func getUserId(c *fiber.Ctx) string {
  token := c.Locals("user").(*jwt.Token)
  return token.Claims.(jwt.MapClaims)["id"].(string)
}

func GetUser(c *fiber.Ctx) error {
  userId := getUserId(c)

  db := database.NewMongoDB()
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

  db := database.NewMongoDB()

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

  db := database.NewMongoDB()
  userId := getUserId(c)

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
  userId := getUserId(c)
  password := c.FormValue("password")

  db := database.NewMongoDB()
  err := db.ChangePassword(database.Password{
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
