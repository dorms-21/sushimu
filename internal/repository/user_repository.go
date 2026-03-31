package repository

import (
	"database/sql"

	"restaurante-go/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, username, password_hash, name, role, active, created_at
		FROM users
		WHERE username = $1
	`, username)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Name, &u.Role, &u.Active, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query(`
		SELECT id, username, password_hash, name, role, active, created_at
		FROM users
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Name, &u.Role, &u.Active, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, u)
	}
	return items, nil
}

func (r *UserRepository) Create(u *models.User) error {
	_, err := r.DB.Exec(`
		INSERT INTO users (username, password_hash, name, role, active)
		VALUES ($1, $2, $3, $4, $5)
	`, u.Username, u.PasswordHash, u.Name, u.Role, u.Active)
	return err
}

func (r *UserRepository) Update(u *models.User) error {
	_, err := r.DB.Exec(`
		UPDATE users
		SET username = $1, name = $2, role = $3, active = $4
		WHERE id = $5
	`, u.Username, u.Name, u.Role, u.Active, u.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}