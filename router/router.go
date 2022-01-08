package router

import (
	"glog/handler"
	"glog/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
  
  // User
  user := api.Group("/user")
  user.Get("/", middleware.Protected(), handler.GetUser)
  user.Post("/", handler.RegisterUser)
  user.Put("/", middleware.Protected(), handler.UpdateUser)
  user.Delete("/", middleware.Protected(), handler.DeleteUser)

  password := user.Group("/password")
  password.Put("/", middleware.Protected(), handler.ChangePassword)
}
