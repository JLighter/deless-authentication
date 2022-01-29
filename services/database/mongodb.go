package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseDriver interface {
  connect() (interface{}, error)
  disconnect(interface{}) error
}

type MongoDB struct {
  ctx context.Context
  client *mongo.Client
}

var lock = &sync.Mutex{}

var instance *MongoDB

func GetMongoDB() *MongoDB {
    lock.Lock()
    defer lock.Unlock()
    if instance == nil {
        ctx := context.Background()
        client, err := connect(ctx)
        if err != nil {
            log.Fatalf("cannot connect to mongodb: %v", err)
        }

        instance = &MongoDB{
          ctx: ctx,
          client: client,
        }
    }
    return instance
}

func connect(ctx context.Context) (*mongo.Client, error) {
  uri := os.Getenv("MONGO_URI")
  if uri == "" {
    return nil, fmt.Errorf("MONGO_URI environment variable is empty")
  }

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
  if err != nil {
    return nil, err
  }

  return client, nil
}

func (m *MongoDB) Ping() error {
  err := m.client.Ping(m.ctx, nil)
  if err != nil {
    return fmt.Errorf("cannot ping mongodb: %v", err)
  }

  return nil
}

