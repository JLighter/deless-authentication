package main

import (
	"glog/middleware"
	"glog/router"
	"glog/services/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.SetupMiddlewares(app)
	router.SetupRoutes(app)

	logger.GetLogger().ApplicationCrashed(app.Listen(":80").Error())
}
