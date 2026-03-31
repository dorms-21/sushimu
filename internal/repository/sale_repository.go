package repository

import (
	"database/sql"

	"restaurante-go/internal/models"
)

type SaleRepository struct {
	DB *sql.DB
}

func NewSaleRepository(db *sql.DB) *SaleRepository {
	return &SaleRepository{DB: db}
}

func (r *SaleRepository) GetAll() ([]models.Sale, error) {
	rows, err := r.DB.Query(`
		SELECT id, order_id, payment_method, amount, status, created_at
		FROM sales
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Sale
	for rows.Next() {
		var s models.Sale
		err := rows.Scan(&s.ID, &s.OrderID, &s.PaymentMethod, &s.Amount, &s.Status, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, s)
	}
	return items, nil
}

func (r *SaleRepository) Create(s *models.Sale) error {
	_, err := r.DB.Exec(`
		INSERT INTO sales (order_id, payment_method, amount, status)
		VALUES ($1, $2, $3, $4)
	`, s.OrderID, s.PaymentMethod, s.Amount, s.Status)
	return err
}

func (r *SaleRepository) Update(s *models.Sale) error {
	_, err := r.DB.Exec(`
		UPDATE sales
		SET order_id = $1, payment_method = $2, amount = $3, status = $4
		WHERE id = $5
	`, s.OrderID, s.PaymentMethod, s.Amount, s.Status, s.ID)
	return err
}

func (r *SaleRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM sales WHERE id = $1`, id)
	return err
}