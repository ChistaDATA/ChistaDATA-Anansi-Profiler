package parsers

import "github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"

// IParser act as a parser for single log line and processes the information to gather info for *stucts.DBPerfInfoRepository
type IParser interface {
	InitParser(DBPerfInfoRepository *stucts.DBPerfInfoRepository, config *stucts.Config)
	Parse(logLine string)
	IsUsable(version string, database string) bool
}
