package models

import "fmt"

const (
	UserNameKey     = "username"
	UserPasswordKey = "password"
	UserDateKey     = "date_created"
)

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Date     string `json:"date_created"`
}

type UserSensitive struct {
	Name string `json:"username"`
	Date string `json:"date_created"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}
