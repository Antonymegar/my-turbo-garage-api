package models

import "myturbogarage/helpers"

type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) SetPassword(password string) {
	hash := helpers.HashPassword(password)
	u.Password = string(hash)
}
