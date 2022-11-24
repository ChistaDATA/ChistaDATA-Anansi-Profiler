package parsers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers/clickhouse"
)

func InitClickHouseParserV1() *Parser {
	return InitParser(
		"0",
		"2.10",
		"clickhouse",
		clickhouse.ParseLogV1,
		[]InfoParserFunc{
			clickhouse.ParseThreadIdV1,
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
