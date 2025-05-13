package models

import (
	"myturbogarage/helpers"
	"time"
)

type User struct {
	ID               string    `json:"id"`
	UserName         string    `json:"userName"`
	Password         string    `json:"password"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	IsMobileVerified bool      `json:"isMobileVerified"`
	IsEmailVerified  bool      `json:"IsEmailVerified"`
	IsActive         bool      `json:"isActive"`
	CreatedAt        time.Time `json:"createdAt"`
	LastLogin        time.Time `json:"lastLogin"`
	CreatedByID      *string    `json:"createdByID"`
	IsAdmin          bool      `json:"isAdmin"`
	ImageUrl         string    `json:"imageUrl"`
}

func (u *User) SetPassword(password string) {
	hash := helpers.HashPassword(password)
	u.Password = string(hash)
}

func (u *User) IsPasswordValid(password string) bool {
	return helpers.ComparePassword([]byte(u.Password), password)
}
