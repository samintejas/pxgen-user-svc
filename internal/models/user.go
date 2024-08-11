package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
