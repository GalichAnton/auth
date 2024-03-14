package model

// Permission ...
type Permission struct {
	Permission string
	RoleName   string
}

// Role ...
type Role struct {
	id   int32  `db:"id"`
	name string `db:"name"`
}
