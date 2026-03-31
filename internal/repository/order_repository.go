package repository

import (
	"database/sql"

	"restaurante-go/internal/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	rows, err := r.DB.Query(`
		SELECT id, table_no, user_id, status, total, created_at
		FROM orders
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var o models.Order
		err := rows.Scan(
			&o.ID,
			&o.TableNo,
			&o.UserID,
			&o.Status,
			&o.Total,
			&o.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}