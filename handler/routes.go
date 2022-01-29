package handler

import (
	"glog/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (h *Handler) Register(app *fiber.App) {
	api := app.Group("/api", logger.New())

	auth := api.Group("/auth")
	auth.Post("/login", h.Login)
  
  user := api.Group("/user")
  user.Post("/", h.RegisterUser)
  user.Get("/", middleware.Protected(), h.GetUser)
  user.Put("/", middleware.Protected(), h.UpdateSelf)
  user.Delete("/", middleware.Protected(), h.DeleteUser)

  password := user.Group("/password")
  password.Put("/", middleware.Protected(), h.ChangePassword)
}
