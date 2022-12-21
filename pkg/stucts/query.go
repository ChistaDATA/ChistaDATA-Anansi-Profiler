package stucts

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strings"
	"sync"
	"time"
)

// Query all the information regarding a query is captured using this struct
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

type PartialQuery struct {
	Timestamp         *time.Time
	QueryId           *string
	ClientHost        *string
	ClientPort        *int
	User              *string
	Query             *string
	Completed         *bool
	ErrorCompleteText *string
	ErrorCode         *string
	ErrorMessage      *string
	ErrorStackTrace   *string
	PeakMemoryUsage   *float64
	ReadRows          *int
	ReadBytes         *float64
	Duration          *float64
	InitialQueryId    *string
	Databases         *types.StringSet
	Tables            *types.StringSet
	ThreadIds         *types.IntSet
}

func (query *Query) Add(extractedQueryInfo PartialQuery) {
	query.Lock.Lock()
	defer query.Lock.Unlock()
	if extractedQueryInfo.Timestamp != nil {
		query.Timestamp = *extractedQueryInfo.Timestamp
	}
	if extractedQueryInfo.QueryId != nil {
		query.QueryId = *extractedQueryInfo.QueryId
	}
	if extractedQueryInfo.ClientHost != nil {
		query.ClientHost = *extractedQueryInfo.ClientHost
	}
	if extractedQueryInfo.ClientPort != nil {
		query.ClientPort = *extractedQueryInfo.ClientPort
	}
	if extractedQueryInfo.User != nil {
		query.User = *extractedQueryInfo.User
	}
	if extractedQueryInfo.Query != nil {
		query.Query = *extractedQueryInfo.Query
	}
	if extractedQueryInfo.Completed != nil {
		query.Completed = *extractedQueryInfo.Completed
	}
	if extractedQueryInfo.ErrorCompleteText != nil {
		query.ErrorCompleteText = *extractedQueryInfo.ErrorCompleteText
	}
	if extractedQueryInfo.ErrorCode != nil {
		query.ErrorCode = *extractedQueryInfo.ErrorCode
	}
	if extractedQueryInfo.ErrorMessage != nil {
		query.ErrorMessage = *extractedQueryInfo.ErrorMessage
	}
	if extractedQueryInfo.ErrorStackTrace != nil {
		query.ErrorStackTrace = *extractedQueryInfo.ErrorStackTrace
	}
	if extractedQueryInfo.PeakMemoryUsage != nil {
		query.PeakMemoryUsage = *extractedQueryInfo.PeakMemoryUsage
	}
	if extractedQueryInfo.ReadRows != nil {

		query.ReadRows = *extractedQueryInfo.ReadRows
	}
	if extractedQueryInfo.ReadBytes != nil {
		query.ReadBytes = *extractedQueryInfo.ReadBytes
	}
	if extractedQueryInfo.Duration != nil {
		query.Duration = *extractedQueryInfo.Duration
	}
	if extractedQueryInfo.InitialQueryId != nil {
		query.InitialQueryId = *extractedQueryInfo.InitialQueryId
	}
	if extractedQueryInfo.Databases != nil {
		query.Databases = *extractedQueryInfo.Databases
	}
	if extractedQueryInfo.Tables != nil {
		query.Tables = *extractedQueryInfo.Tables
	}
	if extractedQueryInfo.ThreadIds != nil {
		if query.ThreadIds == nil {
			query.ThreadIds = *extractedQueryInfo.ThreadIds
		} else {
			for k, _ := range *extractedQueryInfo.ThreadIds {
				query.ThreadIds.Add(k)
			}
		}
	}
}

func (query *Query) GetTransformedQuery() string {
	queryString := strings.TrimSpace(query.Query)
	if len(queryString) > 0 && queryString[len(queryString)-1] == ';' {
		queryString = queryString[:len(queryString)-1]
		queryString = strings.TrimSpace(queryString)
	}
	return queryString
}
