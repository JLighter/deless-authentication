package main

import (
	"context"
	"glog/handler"
	"glog/middleware"
	"glog/services/logger"
	"glog/store"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnvironmentFile(logger *logger.StandardLogger) {
	err := godotenv.Load(".env")
	if err != nil {
    logger.CannotLoadEnvironmentFile(err.Error())
	}
}

func registerHealthz(app *fiber.App, db *mongo.Client, logger *logger.StandardLogger) {
  app.Get("/healthz", func(c *fiber.Ctx) error {
    err := db.Ping(context.Background(), nil)
    if err != nil {
      logger.CannotPingMongoDB(err.Error())
      return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "error", "message": "User database is not available", "data": nil})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Health check is ok!"})
  })
}

func connectToMongo(logger *logger.StandardLogger) *mongo.Client {
  mongoURI := os.Getenv("MONGO_URI")
  if mongoURI == "" {
    logger.CannotGetMongoDBInstance("MONGO_URI environment variable is not set")
    panic("Cannot get mongoDB uri")
  }

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
  if err != nil {
    logger.CannotGetMongoDBInstance("cannot connect to mongoDB")
    panic("Cannot connect to mongoDB")
  }

  return db
}

func main() {
	app := fiber.New()
  logger := logger.GetLogger()

	middleware.SetupMiddlewares(app)

  loadEnvironmentFile(logger)

  db := connectToMongo(logger)

  registerHealthz(app, db, logger)

  passwordStore := store.NewPasswordStore(context.Background(), db)
  userStore := store.NewUserStore(context.Background(), db)

  handler := handler.NewHandler(userStore, passwordStore, logger)
  handler.Register(app)

	logger.ApplicationCrashed(app.Listen(":80").Error())
}
