package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const PASSWORD_COLLECTION = "passwords"

type PasswordStore struct {
  ctx context.Context
  database  *mongo.Database
}

type Password struct {
  Id        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
  UserId    primitive.ObjectID `json:"userid" binding:"required" bson:"userid"`
  Value     string `json:"value" binding:"required" bson:"value"`
}

func NewPasswordStore(ctx context.Context, database *mongo.Database) *PasswordStore {
  return &PasswordStore{ctx, database}
}

func (s *PasswordStore) getPasswordCollection() *mongo.Collection {
  passwords := s.database.Collection(PASSWORD_COLLECTION)

  return passwords
}

func (s *PasswordStore) hashPassword(password string) (string, error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", fmt.Errorf("Error hashing password: %v", err)
  }

  return string(hashedPassword), nil
}

func (s *PasswordStore) ComparePassword(toCompare Password) (bool, error) {
  passwords := s.getPasswordCollection()

  var password Password;

  err := passwords.FindOne(s.ctx, bson.M{"userid": toCompare.UserId}).Decode(&password)
  if err != nil {
    return false, nil
  }

  err = bcrypt.CompareHashAndPassword([]byte(password.Value), []byte(toCompare.Value))
  if err != nil {
    return false, nil
  }

  return true, nil
}

func (s *PasswordStore) ChangePassword(newPassword Password) error {
  passwords := s.getPasswordCollection()
  hashedPasswordValue, err := s.hashPassword(newPassword.Value)

  if err != nil {
    return fmt.Errorf("error hashing password: %v", err)
  }

  newPassword.Value = hashedPasswordValue

  _, err = passwords.ReplaceOne(s.ctx, bson.M{"userid": newPassword.UserId}, newPassword)
	if err != nil {
    return fmt.Errorf("error updating password: %v", err)
	}

	return nil
}

func (s *PasswordStore) InsertPassword(password Password) error {
  passwords := s.getPasswordCollection()
  hashedPasswordValue, err := s.hashPassword(password.Value)

  if err != nil {
    return fmt.Errorf("error hashing password: %v", err)
  }

  password.Value = hashedPasswordValue

  _, err = passwords.InsertOne(s.ctx, password)
  if err != nil {
    return fmt.Errorf("error setting password: %v", err)
  }

  return nil
}
