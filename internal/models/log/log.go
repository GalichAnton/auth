package log

import "time"

// Log ...
type Log struct {
	ID        int64
	Info      Info
	CreatedAt time.Time
}

// Info ...
type Info struct {
	Action   string
	EntityID int64
}
