package handler

import (
	"context"
	"fmt"
	"glog/services/database"
	"glog/services/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

func getUserId(c *fiber.Ctx) (primitive.ObjectID, error) {
  token := c.Locals("user").(*jwt.Token)
  claim := token.Claims.(jwt.MapClaims)
  if claim["id"] == nil {
    return primitive.NilObjectID, fmt.Errorf("no user id in token")
  } else {
    id, err := primitive.ObjectIDFromHex(claim["id"].(string))
    return id, err
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
    logger.GetLogger().CannotGetUser(err.Error())
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

  id, err := db.RegisterUser(database.User{
    Email: c.FormValue("email"),
    Username: c.FormValue("username"),
    Admin: false,
  })

  if err != nil {
    logger.GetLogger().CannotCreateUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  err = db.InsertPassword(database.Password{
    UserId: id,
    Value: c.FormValue("password"),
  })

  if err != nil {
    logger.GetLogger().CannotCreatePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  logger.GetLogger().DidCreateUser(id.Hex())
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
    logger.GetLogger().CannotGetUser(err.Error())
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
    logger.GetLogger().CannotUpdateUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  logger.GetLogger().DidUpdateUser(userId.Hex())
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
    logger.GetLogger().CannotChangePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  logger.GetLogger().DidChangePassword(userId.Hex())
  return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}
