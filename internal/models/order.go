package models

import "time"

type Order struct {
	ID        int
	TableNo   string
	UserID    int
	Status    string
	Total     float64
	CreatedAt time.Time
}