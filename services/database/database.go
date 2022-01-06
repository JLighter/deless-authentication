package database

import (
	"glog/config"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
  "context"
)

type Database struct {
  ctx context.Context
}

func NewDatabase() *Database {
  return &Database{
    ctx: context.Background(),
  }
}

func (d *Database) connect() *redis.Client {
	addr, err := config.Config("REDIS_ADDR")
	if err != nil {
		log.Fatalf("cannot get Redis address: %s", err)
	}

	password, err := config.Config("REDIS_PASSWORD")
	if err != nil {
		log.Fatalf("cannot get Redis password: %s", err)
	}

	database, err := config.Config("REDIS_DATABASE")
	if err != nil {
		log.Fatalf("cannot get Redis database name: %s", err)
	}

  databaseNumber, err := strconv.Atoi(database)
	if err != nil {
		log.Fatalf("cannot convert Database var to int: %s", err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       databaseNumber,
	})
}
