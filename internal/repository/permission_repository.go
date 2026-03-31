package repository

import (
	"database/sql"
)

type PermissionRepository struct {
	DB *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (r *PermissionRepository) GetPermission(role, module string) (map[string]bool, error) {
	row := r.DB.QueryRow(`
		SELECT can_view, can_create, can_edit, can_delete
		FROM role_permissions
		WHERE role = $1 AND module = $2
	`, role, module)

	var canView, canCreate, canEdit, canDelete bool
	err := row.Scan(&canView, &canCreate, &canEdit, &canDelete)
	if err != nil {
		return nil, err
	}

	return map[string]bool{
		"view":   canView,
		"create": canCreate,
		"edit":   canEdit,
		"delete": canDelete,
	}, nil
}