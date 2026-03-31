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

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
	Subtotal  float64
}