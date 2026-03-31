package repository

import (
	"database/sql"

	"restaurante-go/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, price, stock, image_url, active, created_at
		FROM products
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.ImageURL,
			&p.Active,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, description, price, stock, image_url, active, created_at
		FROM products
		WHERE id = $1
	`, id)

	var p models.Product
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
		&p.ImageURL,
		&p.Active,
		&p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) Create(p *models.Product) error {
	_, err := r.DB.Exec(`
		INSERT INTO products (name, description, price, stock, image_url, active)
		VALUES ($1, $2, $3, $4, $5, $6)
	`,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.ImageURL,
		p.Active,
	)
	return err
}

func (r *ProductRepository) Update(p *models.Product) error {
	_, err := r.DB.Exec(`
		UPDATE products
		SET name = $1,
		    description = $2,
		    price = $3,
		    stock = $4,
		    image_url = $5,
		    active = $6
		WHERE id = $7
	`,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.ImageURL,
		p.Active,
		p.ID,
	)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM products WHERE id = $1`, id)
	return err
}