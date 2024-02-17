package user

import (
	"database/sql"
	"time"

	desc "github.com/GalichAnton/auth/pkg/user_v1"
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
	Role     desc.Role
}
