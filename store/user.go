package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DATABASE_NAME = "auth"
const USERS_COLLECTION = "users"
const PASSWORD_COLLECTION = "passwords"

type UserStore struct {
  ctx context.Context
  client *mongo.Client
}

func NewUserStore(ctx context.Context, client *mongo.Client) *UserStore {
  return &UserStore{ctx, client}
}

type User struct {
  Id       primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
  Email    string `json:"email" binding:"required" bson:"email"`
  Admin    bool   `json:"-" bson:"admin"`
}

func (d *UserStore) getUserCollection() *mongo.Collection {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  return users
}

func (d *UserStore) UserExists(email string) bool {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return false
  }

  return true
}

func (d *UserStore) RegisterUser(user User) (primitive.ObjectID, error) {
  users := d.getUserCollection()

  result, err := users.InsertOne(d.ctx, user)

	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error saving user: %v", err)
	}

  return result.InsertedID.(primitive.ObjectID), nil
}

func (d *UserStore) UpdateUser(user User) error {
  users := d.getUserCollection()

  _, err := users.ReplaceOne(d.ctx, bson.M{"_id": user.Id}, user)
	if err != nil {
		return fmt.Errorf("error updating database: %v", err)
	}

	return nil
}

func (d *UserStore) GetUserById(id primitive.ObjectID) (*User, error) {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, nil 
	}

	return &user, nil
}

func (d *UserStore) GetUserByEmail(email string) (*User, error) {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return nil, nil 
  }

  return &user, nil
}

func (d *UserStore) DeleteUser(id string) error {
  users := d.getUserCollection()

  _, err := users.DeleteOne(d.ctx, bson.M{"_id": id})
  if err != nil {
    return fmt.Errorf("error deleting user: %v", err)
  }

  return nil
}
