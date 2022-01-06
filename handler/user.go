package handler

import (
	"context"
  "github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func GetUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}

func RegisterUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}

func UpdateUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}

func DeleteUser(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotImplemented)
}
