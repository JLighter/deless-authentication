package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DATABASE_NAME = "auth"
const USERS_COLLECTION = "users"
const PASSWORD_COLLECTION = "passwords"

type User struct {
  Id       primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
  Username string `json:"username" binding:"required" bson:"username"`
  Email    string `json:"email" binding:"required" bson:"email"`
  Admin    bool   `json:"-" bson:"admin"`
}

func (d *MongoDB) getUserCollection() *mongo.Collection {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  return users
}

func (d *MongoDB) UserExists(email string) bool {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return false
  }

  return true
}

func (d *MongoDB) RegisterUser(user User) (primitive.ObjectID, error) {
  users := d.getUserCollection()

  result, err := users.InsertOne(d.ctx, user)

	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error saving user: %v", err)
	}

  return result.InsertedID.(primitive.ObjectID), nil
}

func (d *MongoDB) UpdateUser(user User) error {
  users := d.getUserCollection()

  _, err := users.ReplaceOne(d.ctx, bson.M{"_id": user.Id}, user)
	if err != nil {
		return fmt.Errorf("Error updating database: %v", err)
	}

	return nil
}

func (d *MongoDB) GetUserById(id primitive.ObjectID) (*User, error) {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, nil 
	}

	return &user, nil
}

func (d *MongoDB) GetUserByEmail(email string) (*User, error) {
  users := d.getUserCollection()

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return nil, nil 
  }

  return &user, nil
}

func (d *MongoDB) DeleteUser(id string) error {
  users := d.getUserCollection()

  _, err := users.DeleteOne(d.ctx, bson.M{"_id": id})
  if err != nil {
    return fmt.Errorf("Error deleting user: %v", err)
  }

  return nil
}
