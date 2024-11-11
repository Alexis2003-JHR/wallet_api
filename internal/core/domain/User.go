package domain

import "time"

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	Password     string `json:"password"`
	LastModified time.Time
}
