package parsers

import (
	"errors"
	"fmt"
)

var parsers = [...]IParser{
	InitClickHouseParserV1(),
}

func GetParser(version string, database string) (IParser, error) {
	for _, parser := range parsers {
		if parser.IsUsable(version, database) {
			return parser, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("A parser for database: %s and version: %s not found", database, version))
}
