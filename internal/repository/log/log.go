package user

import (
	"context"
	"time"

	serviceModel "github.com/GalichAnton/auth/internal/models/log"
	"github.com/GalichAnton/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName    = "logs"
	colID        = "id"
	colAction    = "action"
	colEntityID  = "entity_id"
	colCreatedAt = "created_at"
)

// LogRepository - .
type LogRepository struct {
	db db.Client
}

// NewLogRepository - .
func NewLogRepository(db db.Client) *LogRepository {
	return &LogRepository{db: db}
}

// Create - .
func (l *LogRepository) Create(ctx context.Context, log *serviceModel.Info) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colAction, colEntityID, colCreatedAt).
		Values(log.Action, log.EntityID, time.Now())

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log_repository.Create",
		QueryRaw: query,
	}

	_, err = l.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
