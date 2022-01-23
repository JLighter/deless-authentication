package handler

import (
	"glog/services/database"
	"glog/services/logger"
	"glog/services/token"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	db := database.GetMongoDB()
  user, err := db.GetUserByEmail(email)

  if err != nil {
    logger.GetLogger().CannotComparePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  } else if user == nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

	if ok, err := db.ComparePassword(database.Password{UserId: user.Id, Value: password}); err != nil {
    logger.GetLogger().CannotComparePassword(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := token.GenerateToken(user.Id.Hex(), false)
  if err != nil {
    logger.GetLogger().CannotGenerateToken(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  logger.GetLogger().DidLogin(user.Id.Hex(), c.IP())
	return c.JSON(fiber.Map{"token": token})
}
