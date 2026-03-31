package models

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
	ImageURL    string
	Active      bool
	CreatedAt   time.Time
}