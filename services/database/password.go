package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Password struct {
  Id        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
  UserId    primitive.ObjectID `json:"userid" binding:"required" bson:"userid"`
  Value     string `json:"value" binding:"required" bson:"value"`
}

func (d *MongoDB) getPasswordCollection() *mongo.Collection {
  database := d.client.Database(DATABASE_NAME)
  passwords := database.Collection(PASSWORD_COLLECTION)

  return passwords
}

func (d *MongoDB) ComparePassword(toCompare Password) (bool, error) {
  passwords := d.getPasswordCollection()

  var password Password;

  err := passwords.FindOne(d.ctx, bson.M{"userid": toCompare.UserId}).Decode(&password)
  if err != nil || toCompare.Value != password.Value {
    return false, nil
  }

  return true, nil
}

func (d *MongoDB) ChangePassword(newPassword Password) error {
  passwords := d.getPasswordCollection()

  _, err := passwords.ReplaceOne(d.ctx, bson.M{"userid": newPassword.UserId}, newPassword)
	if err != nil {
    return fmt.Errorf("Error updating password: %v", err)
	}

	return nil
}

func (d *MongoDB) InsertPassword(password Password) error {
  passwords := d.getPasswordCollection()

  _, err := passwords.InsertOne(d.ctx, password)
  if err != nil {
    return fmt.Errorf("Error setting password: %v", err)
  }

  return nil
}
