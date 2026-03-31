package models

import "time"

type Sale struct {
	ID            int
	OrderID       int
	PaymentMethod string
	Amount        float64
	Status        string
	CreatedAt     time.Time
}
