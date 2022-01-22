package handler

import (
	"glog/services/database"
	"glog/services/token"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	db := database.NewMongoDB()
  user, err := db.GetUserByEmail(email)
  if err != nil {
    log.Printf("Error getting user: %v", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if user == nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

	if ok, err := db.ComparePassword(database.Password{UserId: user.Id, Value: password}); err != nil {
		log.Printf("Cannot compare password: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := token.GenerateToken(user.Id, false)

	return c.JSON(fiber.Map{"token": token})
}
