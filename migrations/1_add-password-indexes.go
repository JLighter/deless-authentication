package migrations

import (
	"context"
	"glog/store"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createPasswordUserIDIndex(db *mongo.Database) error {
		opt := options.Index().SetName("userid").SetUnique(true)
		keys := bson.D{{"userid", 1}}
		model := mongo.IndexModel{Keys: keys, Options: opt}
    _, err := db.Collection(store.PASSWORD_COLLECTION).Indexes().CreateOne(context.Background(), model)
    return err
}

func removePasswordUserIDIndex(db *mongo.Database) error {
    _, err := db.Collection(store.PASSWORD_COLLECTION).Indexes().DropOne(context.Background(), "userid")
    return err
}

func init() {
	migrate.Register(func(db *mongo.Database) error {
    err := createPasswordUserIDIndex(db)
		if err != nil {
			return err
		}
	  return nil	
	}, func(db *mongo.Database) error {
    err := removePasswordUserIDIndex(db)
		if err != nil {
			return err
		}
		return nil
	})
}
