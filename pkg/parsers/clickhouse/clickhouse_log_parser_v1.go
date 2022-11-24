package clickhouse

import (
	"errors"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/regexs"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"strconv"
	"time"
)

func ParseLogV1(logLine string) (stucts.ExtractedLog, error) {
	var clickHouseLog stucts.ExtractedLog
	if parts := regexs.ClickHouseLogRegEx.FindStringSubmatch(logLine); len(parts) == 6 {
		var partParsingError error
		clickHouseLog.Timestamp, partParsingError = parseTimestampFromClickHouseLog(parts[1])
		if partParsingError != nil {
			return clickHouseLog, partParsingError
		}
		clickHouseLog.ThreadId, _ = strconv.Atoi(parts[2])
		clickHouseLog.QueryId = parts[3]
		clickHouseLog.LogLevel = parts[4]
		clickHouseLog.Message = parts[5]
		return clickHouseLog, nil
	}
	return clickHouseLog, errors.New("error parsing ClickHouse log, part size not 6")
}

func parseTimestampFromClickHouseLog(s string) (time.Time, error) {
	return time.Parse("2006.01.02 15:04:05.000000", s)
}
