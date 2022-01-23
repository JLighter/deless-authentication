package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

const DATABASE_NAME = "auth"
const USERS_COLLECTION = "users"
const PASSWORD_COLLECTION = "passwords"

type User struct {
  Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
  Admin    bool `json:"-"`
}

type Password struct {
  UserId       string `json:"id" binding:"required"`
  Value     string `json:"value" binding:"required"`
}

func (d *MongoDB) UserExists(email string) (bool, error) {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return false, fmt.Errorf("error checking if user exists: %v", err)
  }

  return true, nil
}

func (d *MongoDB) RegisterUser(user User) error {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  _, err := users.InsertOne(d.ctx, user)

	if err != nil {
		return fmt.Errorf("error saving user: %v", err)
	}

  return nil
}

func (d *MongoDB) ComparePassword(toCompare Password) (bool, error) {
  database := d.client.Database(DATABASE_NAME)
  passwords := database.Collection(PASSWORD_COLLECTION)

  var password Password;

  err := passwords.FindOne(d.ctx, bson.M{"userid": toCompare.UserId}).Decode(&password)
  if err != nil || toCompare.Value != password.Value {
    return false, nil
  }

  return true, nil
}

func (d *MongoDB) ChangePassword(newPassword Password) error {
  database := d.client.Database(DATABASE_NAME)
  passwords := database.Collection(PASSWORD_COLLECTION)

  _, err := passwords.ReplaceOne(d.ctx, bson.M{"UserId": newPassword.UserId}, newPassword)
	if err != nil {
    return fmt.Errorf("Error updating password: %v", err)
	}

	return nil
}

func (d *MongoDB) SetPassword(password Password) error {
  database := d.client.Database(DATABASE_NAME)
  passwords := database.Collection(PASSWORD_COLLECTION)

  _, err := passwords.InsertOne(d.ctx, password)
  if err != nil {
    return fmt.Errorf("Error setting password: %v", err)
  }

  return nil
}

func (d *MongoDB) UpdateUser(user User) error {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  _, err := users.ReplaceOne(d.ctx, bson.M{"id": user.Id}, user)
	if err != nil {
		return fmt.Errorf("Error updating database: %v", err)
	}

	return nil
}

func (d *MongoDB) GetUserById(id string) (*User, error) {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  var user User;
  err := users.FindOne(d.ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, nil 
	}

	return &user, nil
}

func (d *MongoDB) GetUserByEmail(email string) (*User, error) {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  var user User;
  err := users.FindOne(d.ctx, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return nil, nil 
  }

  return &user, nil
}

func (d *MongoDB) DeleteUser(id string) error {
  database := d.client.Database(DATABASE_NAME)
  users := database.Collection(USERS_COLLECTION)

  _, err := users.DeleteOne(d.ctx, bson.M{"Id": id})
  if err != nil {
    return fmt.Errorf("Error deleting user: %v", err)
  }

  return nil
}
