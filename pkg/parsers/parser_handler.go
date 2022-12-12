package parsers

import (
	"errors"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
)

var parsers = [...]IParser{
	InitClickHouseParserV1(),
}

func GetParser(version string, database string, infoCorpus *stucts.InfoCorpus) (IParser, error) {
	for _, parser := range parsers {
		if parser.IsUsable(version, database) {
			parser.SetInfoCorpus(infoCorpus)
			return parser, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("A parser for database: %s and version: %s not found", database, version))
}
