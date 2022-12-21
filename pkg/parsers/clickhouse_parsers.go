package parsers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers/clickhouse"
)

// InitClickHouseParserV1 is Parser object which is a IParser implementation for clickhouse version 0 - 2.10
func InitClickHouseParserV1() *Parser {
	return InitParser(
		"0",
		"2.10",
		"clickhouse",
		clickhouse.ParseLogV1,
		[]InfoParserFunc{
			clickhouse.ParseThreadIdAndTimeV1,
		},
		[]InfoParserFunc{
			clickhouse.ParseMessageWithAccessInfoV1,
			clickhouse.ParseMessageWithQueryV1,
			clickhouse.ParseMessageWithDataInfoV1,
			clickhouse.ParseMessageWithPeakMemoryV1,
			clickhouse.ParseMessageWithErrorInfoV1,
		},
	)
}
