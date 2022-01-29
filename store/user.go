package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS_COLLECTION = "users"

type UserStore struct {
  ctx context.Context
  database *mongo.Database
}

func NewUserStore(ctx context.Context, database *mongo.Database) *UserStore {
  return &UserStore{ctx, database}
}

type User struct {
  Id       primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
  Email    string `json:"email" binding:"required" bson:"email"`
  Admin    bool   `json:"-" bson:"admin"`
}

func (s *UserStore) getUserCollection() *mongo.Collection {
  users := s.database.Collection(USERS_COLLECTION)

  return users
}

func (s *UserStore) UserExists(email string) bool {
  users := s.getUserCollection()

  var user User;
  err := users.FindOne(s.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return false
  }

  return true
}

func (s *UserStore) RegisterUser(user User) (primitive.ObjectID, error) {
  users := s.getUserCollection()

  result, err := users.InsertOne(s.ctx, user)

	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error saving user: %v", err)
	}

  return result.InsertedID.(primitive.ObjectID), nil
}

func (s *UserStore) UpdateUser(user User) error {
  users := s.getUserCollection()

  _, err := users.ReplaceOne(s.ctx, bson.M{"_id": user.Id}, user)
	if err != nil {
		return fmt.Errorf("error updating database: %v", err)
	}

	return nil
}

func (s *UserStore) GetUserById(id primitive.ObjectID) (*User, error) {
  users := s.getUserCollection()

  var user User;
  err := users.FindOne(s.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, nil 
	}

	return &user, nil
}

func (s *UserStore) GetUserByEmail(email string) (*User, error) {
  users := s.getUserCollection()

  var user User;
  err := users.FindOne(s.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return nil, nil 
  }

  return &user, nil
}

func (s *UserStore) DeleteUser(id string) error {
  users := s.getUserCollection()

  _, err := users.DeleteOne(s.ctx, bson.M{"_id": id})
  if err != nil {
    return fmt.Errorf("error deleting user: %v", err)
  }

  return nil
}
