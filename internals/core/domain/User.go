package domain

import "time"

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Address      string
	Email        string
	Age          int
	LastModified time.Time
}
