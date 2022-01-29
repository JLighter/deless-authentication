package handler

import (
	"context"
	"fmt"
	"glog/store"

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

func (h *Handler) GetUser(c *fiber.Ctx) error {
  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  user, err := h.userStore.GetUserById(userId)

  if err != nil {
    h.logger.CannotGetUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if user == nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) RegisterUser(c *fiber.Ctx) error {

  if exists := h.userStore.UserExists(c.FormValue("email")); exists == true {
    return c.SendStatus(fiber.StatusOK)
  }

  id, err := h.userStore.RegisterUser(store.User{
    Email: c.FormValue("email"),
    Admin: false,
  })

  if err != nil {
    h.logger.CannotCreateUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  err = h.passwordStore.InsertPassword(store.Password{
    UserId: id,
    Value: c.FormValue("password"),
  })

  if err != nil {
    h.logger.CannotCreatePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  h.logger.DidCreateUser(id.String())
  return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) UpdateSelf(c *fiber.Ctx) error {

  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  user, err := h.userStore.GetUserById(userId)
  if err != nil {
    h.logger.CannotGetUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if user == nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  newUser := store.User{
    Id: userId,
    Email: c.FormValue("email"),
    Admin: user.Admin,
  }

  err = h.userStore.UpdateUser(newUser)
  if err != nil {
    h.logger.CannotUpdateUser(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  h.logger.DidUpdateUser(userId.String())
  return c.Status(fiber.StatusOK).JSON(newUser)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
  userId, err := getUserId(c)
  if err != nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }
  password := c.FormValue("password")

  err = h.passwordStore.ChangePassword(store.Password{
    UserId: userId,
    Value: password,
  })

  if err != nil {
    h.logger.CannotChangePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  h.logger.DidChangePassword(userId.String())
  return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}
