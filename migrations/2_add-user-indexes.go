package migrations

import (
	"context"
	"glog/store"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUserEmailIndex(db *mongo.Database) error {
		opt := options.Index().SetName("email").SetUnique(true)
		keys := bson.D{{"email", 1}}
		model := mongo.IndexModel{Keys: keys, Options: opt}
    _, err := db.Collection(store.USERS_COLLECTION).Indexes().CreateOne(context.Background(), model)
    return err
}

func removeUserEmailIndex(db *mongo.Database) error {
    _, err := db.Collection(store.USERS_COLLECTION).Indexes().DropOne(context.Background(), "email")
    return err
}

func init() {
	migrate.Register(func(db *mongo.Database) error {
    err := createUserEmailIndex(db)
		if err != nil {
			return err
		}
	  return nil	
	}, func(db *mongo.Database) error {
    err := removeUserEmailIndex(db)
		if err != nil {
			return err
		}
		return nil
	})
}
