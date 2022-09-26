package types

import (
	"strings"
	"time"
)

type Query struct {
	Timestamp         time.Time
	QueryId           string
	ClientHost        string
	ClientPort        int
	User              string
	Query             string
	Completed         bool
	ErrorCompleteText string
	ErrorCode         string
	ErrorMessage      string
	ErrorStackTrace   string
	PeakMemoryUsage   float64
	ReadRows          int
	ReadBytes         float64
	Duration          float64
	InitialQueryId    string
	Databases         StringSet
	Tables            StringSet
	ThreadIds         IntSet
}

func (query *Query) GetTransformedQuery() string {
	queryString := strings.TrimSpace(query.Query)
	if queryString[len(queryString)-1] == ';' {
		queryString = queryString[:len(queryString)-1]
		queryString = strings.TrimSpace(queryString)
	}
	return queryString
}
