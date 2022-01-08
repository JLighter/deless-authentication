package database

import (
	"fmt"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
  Admin    bool
}

type NewUser struct {
	Email    string
	Password string
	Username string
  Admin    bool
}

func userKey(email string) string {
  return fmt.Sprintf("user:%s", email)
}

func userPasswordKey(email string) string {
  return fmt.Sprintf("%s:password", userKey(email))
}

func (d *Database) RegisterUser(user NewUser) (string, error) {
  client, err := d.connect()
  if err != nil {
    return "", fmt.Errorf("Error connecting to database: %s", err)
  }
	defer client.Close()

	pipeline := client.Pipeline()

	pipeline.HMSet(d.ctx, userKey(user.Email), "username", user.Username, "admin", user.Admin)
	pipeline.Set(d.ctx, userPasswordKey(user.Email), user.Password, 0)

	_, err = pipeline.Exec(d.ctx)
	if err != nil {
		return "", fmt.Errorf("error saving user: %s", err)
	}

  return nil
}

// Compare password with the hash stored in the Database
func (d *Database) ComparePassword(email string, password string) (bool, error) {
  client, err := d.connect()
  if err != nil {
    return false, fmt.Errorf("Error connecting to database: %s", err)
  }
  defer client.Close()

  result, err := client.Get(d.ctx, userPasswordKey(email)).Result()
  if err != nil || result != password {
    return false, nil
  }

  return true, nil
}

func (d *Database) ChangePassword(email string, password string) error {
  client, err := d.connect()
  if err != nil {
    return fmt.Errorf("Error connecting to database: %s", err)
  }
	defer client.Close()

	_, err = client.Set(d.ctx, userPasswordKey(email), password, 0).Result()
	if err != nil {
		return fmt.Errorf("Error changing password: %s", err)
	}

	return nil
}

func (d *Database) UpdateUser(user User) error {
	client, err := d.connect()
  if err != nil {
    return fmt.Errorf("Error connecting to database: %s", err)
  }
	defer client.Close()

	_, err = client.HMSet(d.ctx, userKey(user.Email), "username", user.Username).Result()
	if err != nil {
		return fmt.Errorf("Error updating user: %s", err)
	}

	return nil
}

func (d *Database) GetUser(email string) (*User, error) {
	client, err := d.connect()
  if err != nil {
    return nil, fmt.Errorf("Error connecting to database: %s", err)
  }
	defer client.Close()

	result, err := client.HMGet(d.ctx, userKey(email), "username", "admin").Result()
	if err != nil {
		return nil, nil 
	}

	return &User{
		Username: result[0].(string),
		Admin:    result[1].(bool),
		Email:    email,
	}, nil
}
