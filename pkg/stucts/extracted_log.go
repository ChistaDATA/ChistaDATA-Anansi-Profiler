package stucts

import (
	"time"
)

type ExtractedLog struct {
	Timestamp         time.Time
	ThreadId          int
	QueryId           string
	LogLevel          string
	Message           string
	ProcessID         int64
	SessionLineNumber int64
	UserName          string
	DatabaseName      string
	ApplicationName   string
	RemoteHost        string
}
