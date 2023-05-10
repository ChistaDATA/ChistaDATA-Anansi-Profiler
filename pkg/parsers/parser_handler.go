package parsers

import (
	"errors"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
)

var parsers = [...]IParser{
	InitClickHouseParserV1(),
	InitPostgresParserV1(),
}

func GetParser(version string, database string, dBPerfInfoRepository *stucts.DBPerfInfoRepository, config *stucts.Config) (IParser, error) {
	for _, parser := range parsers {
		if parser.IsUsable(version, database) {
			parser.InitParser(dBPerfInfoRepository, config)
			return parser, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("A parser for database: %s and version: %s not found", database, version))
}
