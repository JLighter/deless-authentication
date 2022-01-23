package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

// Hash password using bcrypt
func (d *MongoDB) hashPassword(password string) (string, error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", fmt.Errorf("Error hashing password: %v", err)
  }

  return string(hashedPassword), nil
}

func (d *MongoDB) ComparePassword(toCompare Password) (bool, error) {
  passwords := d.getPasswordCollection()

  var password Password;

  err := passwords.FindOne(d.ctx, bson.M{"userid": toCompare.UserId}).Decode(&password)
  if err != nil {
    return false, nil
  }

  err = bcrypt.CompareHashAndPassword([]byte(password.Value), []byte(toCompare.Value))
  if err != nil {
    return false, nil
  }

  return true, nil
}

func (d *MongoDB) ChangePassword(newPassword Password) error {
  passwords := d.getPasswordCollection()
  hashedPasswordValue, err := d.hashPassword(newPassword.Value)

  if err != nil {
    return fmt.Errorf("error hashing password: %v", err)
  }

  newPassword.Value = hashedPasswordValue

  _, err = passwords.ReplaceOne(d.ctx, bson.M{"userid": newPassword.UserId}, newPassword)
	if err != nil {
    return fmt.Errorf("error updating password: %v", err)
	}

	return nil
}

func (d *MongoDB) InsertPassword(password Password) error {
  passwords := d.getPasswordCollection()
  hashedPasswordValue, err := d.hashPassword(password.Value)

  if err != nil {
    return fmt.Errorf("error hashing password: %v", err)
  }

  password.Value = hashedPasswordValue

  _, err = passwords.InsertOne(d.ctx, password)
  if err != nil {
    return fmt.Errorf("error setting password: %v", err)
  }

  return nil
}
