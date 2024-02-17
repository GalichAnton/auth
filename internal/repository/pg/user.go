package pg

import (
	"context"
	"database/sql"
	"time"

	"github.com/GalichAnton/auth/internal/models/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName    = "users"
	colID        = "id"
	colName      = "name"
	colEmail     = "email"
	colPassword  = "password"
	colRole      = "role"
	colCreatedAt = "created_at"
	colUpdatedAt = "updated_at"
)

// UserRepository - .
type UserRepository struct {
	pool *pgxpool.Pool
}

// NewUserRepository - .
func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

// Create - .
func (u *UserRepository) Create(ctx context.Context, info *user.Info) (int64, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colName, colEmail, colPassword, colRole, colCreatedAt).
		Values(info.Name, info.Email, info.Password, info.Role, time.Now()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = u.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Get - .
func (u *UserRepository) Get(ctx context.Context, id int64) (*user.User, error) {
	builderSelect := sq.Select(colID, colName, colEmail, colPassword, colRole, colCreatedAt, colUpdatedAt).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colID: id})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	var newUser user.User
	var updatedAt sql.NullTime

	row := u.pool.QueryRow(ctx, query, args...)
	err = row.Scan(&newUser.ID, &newUser.Info.Name, &newUser.Info.Email, &newUser.Info.Password, &newUser.Info.Role,
		&newUser.CreatedAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		newUser.UpdatedAt.Time = updatedAt.Time
		newUser.UpdatedAt.Valid = updatedAt.Valid
	}

	return &newUser, nil
}

// Update - .
func (u *UserRepository) Update(ctx context.Context, id int64, info *user.Info) error {
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(colName, info.Name).
		Set(colEmail, info.Email).
		Set(colRole, info.Role).
		Set(colUpdatedAt, time.Now()).
		Where(sq.Eq{colID: id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = u.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Delete - .
func (u *UserRepository) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colID: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = u.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
