package models

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
