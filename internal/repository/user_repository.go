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
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.Name,
		&u.Role,
		&u.Active,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}