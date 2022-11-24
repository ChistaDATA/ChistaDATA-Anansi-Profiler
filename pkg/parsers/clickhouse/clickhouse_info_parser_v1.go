package clickhouse

import (
	"errors"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/regexs"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strconv"
)

func ParseMessageWithQueryV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	query := stucts.PartialQuery{}
	if parts := regexs.LogMessageWithQueryInfoRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 10 {
		query.Query = &parts[8]
		query.ClientHost = &parts[1]
		clientPort, _ := strconv.Atoi(parts[2])
		query.ClientPort = &clientPort
		query.InitialQueryId = &parts[6]
		if len(parts[4]) == 0 {
			parts[4] = "default"
		}
		query.User = &parts[4]
		infoCorpus.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as ExecuteQueryLogMessageWithQueryRegEx")
}

func ParseMessageWithDataInfoV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	query := stucts.PartialQuery{}
	if parts := regexs.LogMessageWithDataRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 7 {
		var partError error
		readRows, _ := strconv.Atoi(parts[1])
		query.ReadRows = &readRows
		readBytes, partError := formattedSizeToBytes(parts[2], parts[4])
		query.ReadBytes = &readBytes
		if partError != nil {
			return partError
		}
		duration, _ := strconv.ParseFloat(parts[5], 64)
		query.Duration = &duration
		completed := true
		query.Completed = &completed
		query.Timestamp = &extractedLog.Timestamp
		infoCorpus.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithDataRegEx")
}

func ParseMessageWithPeakMemoryV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	query := stucts.PartialQuery{}
	if parts := regexs.LogMessageWithPeakMemoryRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 5 {
		var partError error
		peakMemoryUsage, partError := formattedSizeToBytes(parts[2], parts[4])
		query.PeakMemoryUsage = &peakMemoryUsage
		if partError != nil {
			return partError
		}
		infoCorpus.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithPeakMemoryRegEx")
}

func ParseMessageWithErrorInfoV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	query := stucts.PartialQuery{}
	if parts := regexs.LogMessageWithErrorRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 9 {
		query.ErrorCompleteText = &parts[1]
		query.ErrorCode = &parts[3]
		query.ErrorMessage = &parts[4]
		query.ErrorStackTrace = &parts[8]
		infoCorpus.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithErrorRegEx")
}

func ParseMessageWithAccessInfoV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	databases := types.InitStringSet()
	tables := types.InitStringSet()
	query := stucts.PartialQuery{Databases: &databases, Tables: &tables}
	if parts := regexs.LogMessageWithAccessInfoRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 3 {
		query.Databases.Add(parts[1])
		query.Tables.Add(parts[2])
		infoCorpus.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithAccessInfoRegEx")
}

func ParseThreadIdV1(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	threadIds := types.InitIntSet()
	query := stucts.PartialQuery{ThreadIds: &threadIds}
	query.ThreadIds.Add(extractedLog.ThreadId)
	infoCorpus.Queries.Add(query, extractedLog)
	return nil
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
