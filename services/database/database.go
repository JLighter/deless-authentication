package database

import (
	"fmt"
	"glog/config"

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

func (d *Database) connect() (*redis.ClusterClient, error) {
	addr, err := config.Config("REDIS_ADDR")
	if err != nil {
		return nil, fmt.Errorf("cannot get Redis address: %s", err)
	}

  clusterSlots := func(ctx context.Context) ([]redis.ClusterSlot, error) {
    slots := []redis.ClusterSlot{
      // First node with 1 master and 1 slave.
      {
        Start: 0,
        End:   16384,
        Nodes: []redis.ClusterNode{{
          Addr: addr, // master
        }},
      },
    }
    return slots, nil
  }

	return redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: clusterSlots,
    ReadOnly: true,
	}), nil
}
