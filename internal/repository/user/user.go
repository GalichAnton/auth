package user

import (
	"context"
	"time"

	serviceModel "github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository/user/converter"
	modelRepo "github.com/GalichAnton/auth/internal/repository/user/model"
	"github.com/GalichAnton/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

const (
	tableName    = "users"
	colID        = "id"
	colName      = "name"
	colEmail     = "email"
	colPassword  = "password"
	colRole      = "role_id"
	colCreatedAt = "created_at"
	colUpdatedAt = "updated_at"
)

// Repository - .
type Repository struct {
	db db.Client
}

// NewUserRepository - .
func NewUserRepository(db db.Client) *Repository {
	return &Repository{db: db}
}

// Create - .
func (r *Repository) Create(ctx context.Context, info *serviceModel.Info) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	info.Password = string(hashedPassword)

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colName, colEmail, colPassword, colRole, colCreatedAt).
		Values(info.Name, info.Email, info.Password, info.Role, time.Now()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Get - .
func (r *Repository) Get(ctx context.Context, filter modelRepo.Filter) (*serviceModel.User, error) {
	builderSelect := sq.Select(colID, colName, colEmail, colPassword, colRole, colCreatedAt, colUpdatedAt).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	if filter.ID != nil {
		builderSelect = builderSelect.Where(sq.Eq{colID: *filter.ID})
	}

	if filter.Email != nil {
		builderSelect = builderSelect.Where(sq.Eq{colEmail: *filter.Email})
	}

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var newUser modelRepo.User

	err = r.db.DB().ScanOneContext(ctx, &newUser, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToServiceUser(&newUser), nil
}

// Update - .
func (r *Repository) Update(ctx context.Context, id int64, info *serviceModel.Info) error {
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

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// Delete - .
func (r *Repository) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colID: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
