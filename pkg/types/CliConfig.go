package types

import "os"

type ReportType string

var (
	ReportTypeMarkDown = ReportType("md")
	ReportTypeText     = ReportType("text")
)

type CliConfig struct {
	Files         []string
	ReportType    ReportType
	TopQueryCount int
}

func InitializeCliConfig() CliConfig {
	cliConfig := CliConfig{
		ReportType:    ReportTypeText,
		TopQueryCount: 10,
	}

	args := os.Args[1:]

	if len(args) == 1 {
		cliConfig.Files = append(cliConfig.Files, args[0])
		cliConfig.Files = append(cliConfig.Files, "/Users/chistadata/sandbox/ChistaDATA-Profiler-for-ClickHouse/resources/logs/clickhouser-server-sample-1.log")
	}

	return cliConfig
}
