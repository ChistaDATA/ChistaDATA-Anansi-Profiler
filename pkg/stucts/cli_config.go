package stucts

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

// Default config values
const (
	TopQueryCountDefault     = 10
	ReportTypeText           = "text"
	ReportTypeMD             = "md"
	ReportTypeDefault        = ReportTypeText
	MinimumQueryCountDefault = 1
	ClickHouseDatabase       = "clickhouse"
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
	LogLevelPanic            = "panic"
	LogLevelFatal            = "fatal"
	LogLevelError            = "error"
	LogLevelWarn             = "warn"
	LogLevelInfo             = "info"
	LogLevelDebug            = "debug"
	LogLevelTrace            = "trace"
)

var SortFieldOperations = [...]string{SortFieldOperationSum, SortFieldOperationMin, SortFieldOperationMax, SortFieldOperationAvg, SortFieldOperationPer95, SortFieldOperationStdDev, SortFieldOperationMedian}
var SortFields = [...]string{SortFieldExecTime, SortFieldRowsRead, SortFieldBytesRead, SortFieldPeakMemory, SortFieldQPS, SortFieldQueryCount}
var SortOrders = [...]string{SortOrderAsc, SortOrderDesc}

var LogLevels = map[string]uint32{LogLevelPanic: 0, LogLevelFatal: 1, LogLevelError: 2, LogLevelWarn: 3, LogLevelInfo: 4, LogLevelDebug: 5, LogLevelTrace: 6}

// ReportTypes List of supported report types
var ReportTypes = [...]string{ReportTypeText, ReportTypeMD}

var DatabaseNames = [...]string{ClickHouseDatabase}

// TODO : construct config from cli-config

type CliConfig struct {
	TopQueryCount         int      `short:"n" help:"Count of queries for top x table" default:"10"`
	ReportType            string   `short:"r" help:"Report type to be generated, types: md, text" default:"text"`
	FilePaths             []string `arg:"" optional:"" help:"Paths of log files" type:"existingfile"`
	MinimumQueryCallCount int      `short:"c" help:"Minimum no of query calls needed" default:"1"`
	DatabaseName          string   `help:"database type" default:"clickhouse"`
	DatabaseVersion       string   `help:"database version" default:"0"` //TODO make this a supported stable version
	SortField             string   `help:"Sort queries by the given field, possible values: ExecTime, RowsRead, BytesRead, PeakMemory, QPS, QueryCount" default:"ExecTime"`
	SortFieldOperation    string   `help:"Sort queries by the given operation on field, possible values: sum, min, max, avg, per95, stdDev, median" default:"max"`
	SortOrder             string   `help:"Sort order, possible values: asc, desc" default:"desc"`
	LogLevel              string   `help:"log level, possible values: panic, fatal, error, warn, info, debug, trace" default:"error"`
	S3AccessKeyID         string   `name:"s3-access-key-id"`
	S3SecretAccessKey     string   `name:"s3-secret-access-key"`
	S3SessionToken        string   `name:"s3-session-token"`
	S3Region              string   `name:"s3-region"`
	S3FileLocations       []string `name:"s3-object-urls"`
}

func InitializeCliConfig() *CliConfig {
	cliConfig := CliConfig{}
	kong.Parse(&cliConfig)
	cliConfig.validateCliConfig()
	return &cliConfig
}

// validateCliConfig Validating CliConfig inputs from user
func (cliConfig *CliConfig) validateCliConfig() {

	valid := false
	for _, s := range ReportTypes {
		if s == cliConfig.ReportType {
			valid = true
			break
		}
	}
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

	valid = false
	for _, s := range DatabaseNames {
		if s == cliConfig.DatabaseName {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid Database name, Falling back to default")
		cliConfig.DatabaseName = ClickHouseDatabase
	}

	valid = false
	for _, s := range SortFields {
		if s == cliConfig.SortField {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid SortField name, Falling back to default")
		cliConfig.SortField = SortFieldExecTime
	}

	valid = false
	for _, s := range SortFieldOperations {
		if s == cliConfig.SortFieldOperation {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid SortFieldOperation name, Falling back to default")
		cliConfig.SortFieldOperation = SortFieldOperationMax
	}

	valid = false
	for _, s := range SortOrders {
		if s == cliConfig.SortOrder {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid SortOrder name, Falling back to default")
		cliConfig.SortOrder = SortOrderDesc
	}

	valid = false
	for s, _ := range LogLevels {
		if s == cliConfig.LogLevel {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid LogLevel name, Falling back to default")
		cliConfig.LogLevel = LogLevelError
	}
}
