package model

import (
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID        int64        `db:"id"`
	Info      Info         `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Info ...
type Info struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     int32  `db:"role_id"`
}

// Filter ...
type Filter struct {
	ID    *int64
	Email *string
}
