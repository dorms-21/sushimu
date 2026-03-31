package models

import "time"

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Name         string
	Role         string
	Active       bool
	CreatedAt    time.Time
}