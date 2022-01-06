package handler

import (
	"github.com/gofiber/fiber/v2"
  "glog/services/token"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Generate encoded token and send it as response.
  token := token.GenerateToken("John Doe", false)

	return c.JSON(fiber.Map{"token": token})
}
