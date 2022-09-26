package types

import (
	"time"
)

type ClickHouseLog struct {
	Timestamp time.Time
	ThreadId  int
	QueryId   string //uuid.UUID
	LogLevel  string
	Message   string
}
