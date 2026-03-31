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

	var items []models.Order
	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.ID, &o.TableNo, &o.UserID, &o.Status, &o.Total, &o.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, o)
	}
	return items, nil
}

func (r *OrderRepository) Create(o *models.Order) error {
	_, err := r.DB.Exec(`
		INSERT INTO orders (table_no, user_id, status, total)
		VALUES ($1, $2, $3, $4)
	`, o.TableNo, o.UserID, o.Status, o.Total)
	return err
}

func (r *OrderRepository) Update(o *models.Order) error {
	_, err := r.DB.Exec(`
		UPDATE orders
		SET table_no = $1, user_id = $2, status = $3, total = $4
		WHERE id = $5
	`, o.TableNo, o.UserID, o.Status, o.Total, o.ID)
	return err
}

func (r *OrderRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM orders WHERE id = $1`, id)
	return err
}