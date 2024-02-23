package user

import (
	"database/sql"
	"time"
)

// User - .
type User struct {
	ID        int64
	Info      Info
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// Info - .
type Info struct {
	Name     string
	Email    string
	Password string
	Role     int32
}

// ToCreate - .
type ToCreate struct {
	Info
	PasswordConfirm string
}
