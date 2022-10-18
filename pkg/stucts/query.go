package stucts

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strings"
	"sync"
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
	Databases         types.StringSet
	Tables            types.StringSet
	ThreadIds         types.IntSet
	Lock              sync.Mutex
}

func (query *Query) GetTransformedQuery() string {
	queryString := strings.TrimSpace(query.Query)
	if len(queryString) > 0 && queryString[len(queryString)-1] == ';' {
		queryString = queryString[:len(queryString)-1]
		queryString = strings.TrimSpace(queryString)
	}
	return queryString
}
