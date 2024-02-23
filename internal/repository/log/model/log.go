package model

import "time"

// Log - .
type Log struct {
	ID        int64     `db:"id"`
	Info      LogInfo   `db:""`
	CreatedAt time.Time `db:"created_at"`
}

// LogInfo - .
type LogInfo struct {
	Action   string `db:"action"`
	EntityID int64  `db:"entity_id"`
}
