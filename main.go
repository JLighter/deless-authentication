package main

import (
	"glog/middleware"
	"glog/router"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.SetupMiddlewares(app)
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":80"))
}
