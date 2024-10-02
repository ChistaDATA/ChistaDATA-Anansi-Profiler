package stucts

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

var TempFolder string

// Default config values
const (
	TopQueryCountDefault     = 10
	MinimumQueryCountDefault = 1

	ReportTypeText    = "text"
	ReportTypeMD      = "md"
	ReportTypeDefault = ReportTypeText

	ClickHouseDatabase = "clickhouse"
	PostgresDatabase   = "postgres"

	SortFieldExecTime        = "ExecTime"
	SortFieldRowsRead        = "RowsRead"
	SortFieldBytesRead       = "BytesRead"
	SortFieldPeakMemory      = "PeakMemory"
	SortFieldQPS             = "QPS"
	SortFieldQueryCount      = "QueryCount"
	SortOrderAsc             = "asc"
	SortOrderDesc            = "desc"
	SortFieldOperationSum    = "sum"
	SortFieldOperationMin    = "min"
	SortFieldOperationMax    = "max"
	SortFieldOperationAvg    = "avg"
	SortFieldOperationPer95  = "per95"
	SortFieldOperationStdDev = "stdDev"
	SortFieldOperationMedian = "median"

	LogLevelPanic = "panic"
	LogLevelFatal = "fatal"
	LogLevelError = "error"
	LogLevelWarn  = "warn"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
	LogLevelTrace = "trace"

	SelectQuery = "select"
	CreateQuery = "create"
	UpdateQuery = "update"
	DeleteQuery = "delete"
	InsertQuery = "insert"
)

var SortFieldOperations = []string{SortFieldOperationSum, SortFieldOperationMin, SortFieldOperationMax, SortFieldOperationAvg, SortFieldOperationPer95, SortFieldOperationStdDev, SortFieldOperationMedian}
var SortFields = []string{SortFieldExecTime, SortFieldRowsRead, SortFieldBytesRead, SortFieldPeakMemory, SortFieldQPS, SortFieldQueryCount}
var SortOrders = []string{SortOrderAsc, SortOrderDesc}

var LogLevels = map[string]uint32{LogLevelPanic: 0, LogLevelFatal: 1, LogLevelError: 2, LogLevelWarn: 3, LogLevelInfo: 4, LogLevelDebug: 5, LogLevelTrace: 6}

// ReportTypes List of supported report types
var ReportTypes = []string{ReportTypeText, ReportTypeMD}

var DatabaseNames = []string{ClickHouseDatabase, PostgresDatabase}

var DiscardQueries = []string{SelectQuery, CreateQuery, UpdateQuery, DeleteQuery, InsertQuery}

// TODO : construct config from cli-config

type Config struct {
	TopQueryCount         int      `short:"n" help:"Count of queries for top x table" default:"10"`
	ReportType            string   `short:"r" help:"Report type to be generated, types: md, text" default:"text"`
	FilePaths             []string `arg:"" optional:"" help:"Paths of log files" type:"existingfile"`
	MinimumQueryCallCount int      `short:"c" help:"Minimum no of query calls needed" default:"1"`
	DatabaseType          string   `help:"Which database? Possible values: clickhouse, postgres" default:"clickhouse"`
	DatabaseVersion       string   `help:"Database version" default:"0"` //TODO make this a supported stable version
	SortField             string   `help:"Sort queries by the given field, possible values: ExecTime, RowsRead, BytesRead, PeakMemory, QPS, QueryCount" default:"ExecTime"`
	SortFieldOperation    string   `help:"Sort queries by the given operation on field, possible values: sum, min, max, avg, per95, stdDev, median" default:"max"`
	SortOrder             string   `help:"Sort order, possible values: asc, desc" default:"desc"`
	LogLevel              string   `help:"Log level, possible values: panic, fatal, error, warn, info, debug, trace" default:"error"`
	LogPrefix             string   `help:"Prefix of log" default:""`
	DiscardQueries        []string `help:"It will consider all the query types by default but type of queries can be discarded, possible values: select, update, delete, insert" default:""`
	S3Config              S3Config `embed:"s3" prefix:"s3-"`
}

func InitializeCliConfig() *Config {
	cliConfig := Config{}
	kong.Parse(&cliConfig)
	cliConfig.validateCliConfig()
	return &cliConfig
}

// validateCliConfig Validating Config inputs from user
func (cliConfig *Config) validateCliConfig() {
	valid := cliConfig.isArgumentListValid(cliConfig.ReportType, ReportTypes)
	if !valid {
		log.Warningln("Invalid Report type, Falling back to default")
		cliConfig.ReportType = ReportTypeDefault
	}

	valid = false
	if cliConfig.TopQueryCount <= 0 {
		log.Warningln("Invalid Top Query Count, Falling back to default")
		cliConfig.TopQueryCount = TopQueryCountDefault
	}

	if cliConfig.MinimumQueryCallCount <= 0 {
		log.Warningln("Invalid Minimum Query Count, Falling back to default")
		cliConfig.MinimumQueryCallCount = MinimumQueryCountDefault
	}

	valid = cliConfig.isArgumentListValid(cliConfig.DatabaseType, DatabaseNames)
	if !valid {
		log.Warningln("Invalid Database name, Falling back to default")
		cliConfig.DatabaseType = ClickHouseDatabase
	}

	valid = cliConfig.isArgumentListValid(cliConfig.SortField, SortFields)
	if !valid {
		log.Warningln("Invalid SortField name, Falling back to default")
		cliConfig.SortField = SortFieldExecTime
	}

	valid = cliConfig.isArgumentListValid(cliConfig.SortFieldOperation, SortFieldOperations)
	if !valid {
		log.Warningln("Invalid SortFieldOperation name, Falling back to default")
		cliConfig.SortFieldOperation = SortFieldOperationMax
	}

	valid = cliConfig.isArgumentListValid(cliConfig.SortOrder, SortOrders)
	if !valid {
		log.Warningln("Invalid SortOrder name, Falling back to default")
		cliConfig.SortOrder = SortOrderDesc
	}

	valid = false
	for s := range LogLevels {
		if s == cliConfig.LogLevel {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid LogLevel name, Falling back to default")
		cliConfig.LogLevel = LogLevelError
	}

	valid = isSubsetStringArray(cliConfig.DiscardQueries, DiscardQueries)
	if !valid {
		log.Warningln("Invalid DiscardQueries name, Falling back to default")
		cliConfig.DiscardQueries = []string{CreateQuery, UpdateQuery, DeleteQuery, InsertQuery}
	}
}

func (cliConfig *Config) isArgumentListValid(value string, validOptions []string) bool {
	valid := false
	for _, s := range validOptions {
		if s == value {
			valid = true
			break
		}
	}
	return valid
}

func isSubsetStringArray(sub []string, main []string) bool {
	isSubset := true
	for i := 0; i < len(sub); i++ {
		found := false
		for j := 0; j < len(main); j++ {
			if sub[i] == main[j] {
				found = true
				break
			}
		}
		if !found {
			isSubset = false
			break
		}
	}
	return isSubset
}
