package database

import (
	"fmt"
	"glog/config"
	"strconv"

	"context"

	"github.com/go-redis/redis/v8"
)

type Database struct {
  ctx context.Context
}

func NewDatabase() *Database {
  return &Database{
    ctx: context.Background(),
  }
}

func (d *Database) connect() (*redis.Client, error) {
	addr, err := config.Config("REDIS_ADDR")
	if err != nil {
		return nil, fmt.Errorf("cannot get Redis address: %s", err)
	}

	password, err := config.Config("REDIS_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("cannot get Redis password: %s", err)
	}

	database, err := config.Config("REDIS_DATABASE")
	if err != nil {
		return nil, fmt.Errorf("cannot get Redis database name: %s", err)
	}

  databaseNumber, err := strconv.Atoi(database)
	if err != nil {
		return nil, fmt.Errorf("cannot convert Database var to int: %s", err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       databaseNumber,
	}), nil
}
