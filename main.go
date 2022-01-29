package main

import (
	"glog/middleware"
	"glog/router"
	"glog/services/logger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	middleware.SetupMiddlewares(app)
	router.SetupRoutes(app)

	err := godotenv.Load(".env")
	if err != nil {
    log.Print("cannot load .env file, fallback to env variables")
	}

	logger.GetLogger().ApplicationCrashed(app.Listen(":80").Error())
}
