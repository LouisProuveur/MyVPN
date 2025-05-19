package models

import (
	"encoding/json"
	"os"
)

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadUsers() ([]User, error) {
	data, err := os.ReadFile("/app/config/users.json")
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func AddNewUser(username string, pwHash string) error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}

	var user User
	user.Username = username
	user.PasswordHash = pwHash

	users = append(users, user)
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("/app/config/users.json", data, 0600)
}
