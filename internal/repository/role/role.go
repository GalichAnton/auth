package role

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/role"
	"github.com/GalichAnton/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"
)

type Repository struct {
	db db.Client
}

// NewRoleRepository - .
func NewRoleRepository(db db.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllRolePermissions(ctx context.Context) ([]role.Permission, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	selectBuilder := psql.Select("r.id, rp.permission").
		From("role_permissions rp").
		Join("roles r ON r.id = rp.role_id")

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "role_repository.GetAllRolePermissions",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []role.Permission

	for rows.Next() {
		var permission role.Permission

		err := rows.Scan(&permission.RoleID, &permission.Permission)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}
