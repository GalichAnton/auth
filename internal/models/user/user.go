package user

import (
	"database/sql"
	"time"

	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// UserInfo - .
type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     desc.Role
}
