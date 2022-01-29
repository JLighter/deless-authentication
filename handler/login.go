package handler

import (
	"glog/services/token"
	"glog/store"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

  user, err := h.userStore.GetUserByEmail(email)

  if err != nil {
    h.logger.CannotComparePassword(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  } else if user == nil {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

	if ok, err := h.passwordStore.ComparePassword(store.Password{UserId: user.Id, Value: password}); err != nil {
    h.logger.CannotComparePassword(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := token.GenerateToken(user.Id.Hex(), false)
  if err != nil {
    h.logger.CannotGenerateToken(err.Error())
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  h.logger.DidLogin(user.Id.Hex(), c.IP())
	return c.JSON(fiber.Map{"token": token})
}
