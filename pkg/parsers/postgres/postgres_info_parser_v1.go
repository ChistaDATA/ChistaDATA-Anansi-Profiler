package postgres

import (
	"errors"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strconv"
)

func ParseMessageWithNewQueryV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithNewQueryRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 2 {
		dBPerfInfoRepository.CurrentQuery = &stucts.Query{QueryId: extractedLog.QueryId}
		processIDs := types.InitIntSet()
		processIDs.Add(extractedLog.ThreadId)
		completed := true
		query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId, Timestamp: &extractedLog.Timestamp, ProcessIDs: &processIDs, Completed: &completed}
		query.Query = &parts[1]
		dBPerfInfoRepository.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithNewQueryRegEx")
}

func ParseMessageWithQueryDurationV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithDurationRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 3 {
		if dBPerfInfoRepository.CurrentQuery != nil {
			query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId}
			query.Duration = parseDuration(parts[1], parts[2])
			dBPerfInfoRepository.Queries.Add(query, extractedLog)
			dBPerfInfoRepository.CurrentQuery = nil
		}
		return nil
	}
	return errors.New("error parsing message as LogMessageWithDurationRegEx")
}

func ParseLogMessageWithErrorRegExV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithErrorRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 2 {
		if dBPerfInfoRepository.CurrentQuery != nil {
			completed := false
			query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId, Completed: &completed}
			dBPerfInfoRepository.Queries.Add(query, extractedLog)
		}
		return nil
	}
	return errors.New("error parsing message as LogMessageWithErrorRegEx")
}

//func ParseLogMessageWithEndQueryRegExV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
//	if parts := PostgresLogMessageWithEndQueryRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 2 {
//		dBPerfInfoRepository.CurrentQuery = nil
//		return nil
//	}
//	return errors.New("error parsing message as LogMessageWithEndQueryRegEx")
//}

func parseDuration(s string, unit string) *float64 {
	sizeInFloat, _ := strconv.ParseFloat(s, 64)
	if unit == "ms" {
		sizeInFloat /= 1000
	}
	return &sizeInFloat
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
