package handler

import (
	"glog/services/database"
	"glog/services/token"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

  db := database.NewDatabase()
  
  if ok, err := db.ComparePassword(email, password); err != nil {
    return c.SendStatus(fiber.StatusServiceUnavailable)
  } else if !ok {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

	// Generate encoded token and send it as response.
  token := token.GenerateToken(email, false)

	return c.JSON(fiber.Map{"token": token})
}
