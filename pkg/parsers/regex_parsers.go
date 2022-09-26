package parsers

import (
	"errors"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strconv"
	"time"
)

func ParseClickHouseLog(logLine string) (types.ClickHouseLog, error) {
	var clickHouseLog types.ClickHouseLog
	if parts := ClickHouseLogRegEx.FindStringSubmatch(logLine); len(parts) == 6 {
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

func ParseMessageWithQuery(message string, query *types.Query) error {
	if parts := LogMessageWithQueryInfoRegEx.FindStringSubmatch(message); len(parts) == 10 {
		query.Query = parts[8]
		query.ClientHost = parts[1]
		query.ClientPort, _ = strconv.Atoi(parts[2])
		query.InitialQueryId = parts[6]
		query.User = parts[4]
		if len(query.User) == 0 {
			query.User = "default"
		}
		return nil
	}
	return errors.New("error parsing message as ExecuteQueryLogMessageWithQueryRegEx")
}

func ParseMessageWithDataInfo(message string, query *types.Query) error {
	if parts := LogMessageWithDataRegEx.FindStringSubmatch(message); len(parts) == 7 {
		var partError error
		query.ReadRows, _ = strconv.Atoi(parts[1])
		query.ReadBytes, partError = formattedSizeToBytes(parts[2], parts[4])
		if partError != nil {
			return partError
		}
		query.Duration, _ = strconv.ParseFloat(parts[5], 64)
		query.Completed = true
		return nil
	}
	return errors.New("error parsing message as LogMessageWithDataRegEx")
}

func ParseMessageWithPeakMemory(message string, query *types.Query) error {
	if parts := LogMessageWithPeakMemoryRegEx.FindStringSubmatch(message); len(parts) == 5 {
		var partError error
		query.PeakMemoryUsage, partError = formattedSizeToBytes(parts[2], parts[4])
		if partError != nil {
			return partError
		}
		return nil
	}
	return errors.New("error parsing message as LogMessageWithPeakMemoryRegEx")
}

func ParseMessageWithErrorInfo(message string, query *types.Query) error {
	if parts := LogMessageWithErrorRegEx.FindStringSubmatch(message); len(parts) == 9 {
		query.ErrorCompleteText = parts[1]
		query.ErrorCode = parts[3]
		query.ErrorMessage = parts[4]
		query.ErrorStackTrace = parts[8]
		return nil
	}
	return errors.New("error parsing message as LogMessageWithErrorRegEx")
}

func ParseMessageWithAccessInfo(message string, query *types.Query) error {
	if parts := LogMessageWithAccessInfoRegEx.FindStringSubmatch(message); len(parts) == 3 {
		query.Databases.Add(parts[1])
		query.Tables.Add(parts[2])
		return nil
	}
	return errors.New("error parsing message as LogMessageWithAccessInfoRegEx")
}

func formattedSizeToBytes(size string, unit string) (float64, error) {
	sizeInFloat, _ := strconv.ParseFloat(size, 64)
	if unit == "B" {
		return sizeInFloat, nil
	} else if unit == "KiB" {
		return sizeInFloat * 1024, nil
	} else if unit == "MiB" {
		return sizeInFloat * 1024 * 1024, nil
	} else if unit == "GiB" {
		return sizeInFloat * 1024 * 1024 * 1024, nil
	}
	return 0, errors.New("size format " + unit + " not supported")
}

// parseTimestampFromClickHouseLog layout is 2022.09.08 05:09:25.715819
func parseTimestampFromClickHouseLog(s string) (time.Time, error) {
	return time.Parse("2006.01.02 15:04:05.000000", s)
}
