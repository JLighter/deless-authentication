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

	api.Get("/test", middleware.Protected(), handler.Test)
  
  // User
  user := api.Group("/user")
  user.Get("/:id", middleware.Protected(), handler.GetUser)
  user.Post("/", handler.RegisterUser)
  user.Put("/:id", middleware.Protected(), handler.UpdateUser)
  user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
}
