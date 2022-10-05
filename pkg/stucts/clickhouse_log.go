package stucts

import (
	"time"
)

type ClickHouseLog struct {
	Timestamp time.Time
	ThreadId  int
	QueryId   string
	LogLevel  string
	Message   string
}
