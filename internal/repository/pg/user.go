package pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/GalichAnton/auth/internal/models/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName = "users"
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
		Columns("name", "email", "password", "role").
		Values(info.Name, info.Email, info.Password, info.Role).
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
	builderSelect := sq.Select("id", "name", "email", "password", "role", "created_at", "updated_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

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
	}

	return &newUser, nil
}

// Update - .
func (u *UserRepository) Update(ctx context.Context, id int64, info *user.Info) error {
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("name", info.Name).
		Set("email", info.Email).
		Set("role", info.Role).
		Where(sq.Eq{"id": id})

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
		Where(sq.Eq{"id": id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	_, err = u.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
