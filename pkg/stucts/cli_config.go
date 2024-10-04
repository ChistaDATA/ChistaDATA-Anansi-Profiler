package stucts

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

var TempFolder string

// Default config values
const (
	ReportTypeText = "text"
	ReportTypeMD   = "md"
	ReportTypeHTML = "html"

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
)

type Config struct {
	FilePaths             []string `arg:"" optional:"" help:"Paths of log files" type:"existingfile"`
	TopQueryCount         int      `short:"n" help:"Count of queries for top x table" default:"10"`
	ReportType            string   `short:"r" help:"Report type to be generated. Possible values: ${enum}" default:"text" enum:"md,text,html"`
	MinimumQueryCallCount int      `short:"c" help:"Minimum no of query calls needed" default:"1"`
	DatabaseType          string   `help:"Database type where the log file was generated. Possible values: ${enum}" default:"clickhouse" enum:"clickhouse,postgres"`
	DatabaseVersion       string   `help:"Database version" default:"0"` //TODO make this a supported stable version
	SortField             string   `help:"Sort queries by the given field. Possible values: ${enum}" default:"ExecTime" enum:"ExecTime,RowsRead,BytesRead,PeakMemory,QPS,QueryCount"`
	SortFieldOperation    string   `help:"Sort queries by the given operation on field. Possible values: ${enum}" default:"max" enum:"sum,min,max,avg,per95,stdDev,median"`
	SortOrder             string   `help:"Sort order. Possible values: ${enum}" default:"desc" enum:"asc,desc"`
	LogLevel              string   `help:"Log level. Possible values: ${enum}" default:"error" enum:"panic,fatal,error,warn,info,debug,trace"`
	LogPrefix             string   `help:"Prefix of log" default:""`
	DiscardQueries        []string `help:"It will consider all the query types by default but type of queries can be discarded. Possible values: ${enum}" default:"" enum:"select,create,update,delete,insert"`
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
	if cliConfig.TopQueryCount <= 0 {
		log.Fatalln("Invalid Top Query Count, should be greater than 0")
	}

	if cliConfig.MinimumQueryCallCount <= 0 {
		log.Fatalln("Invalid Minimum Query Count, should be greater than 0")
	}
}
