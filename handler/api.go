package handler

import (
	"glog/services/database"
	"glog/services/logger"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck handle api health check
func HealthCheck(c *fiber.Ctx) error {
  err := database.GetMongoDB().Ping()
  if err != nil {
    logger.GetLogger().CannotPingMongoDB(err.Error())
    return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "error", "message": "User database is not available", "data": nil})
  }
  return c.JSON(fiber.Map{"status": "success", "message": "Health check is ok!"})
}
