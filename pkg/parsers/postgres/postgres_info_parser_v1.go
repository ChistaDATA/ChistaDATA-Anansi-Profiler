package postgres

import (
	"errors"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"strconv"
)

func ParseMessageWithNewQueryV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithNewQueryRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 2 {
		query := getNewQuery(extractedLog, dBPerfInfoRepository)
		query.Query = &parts[1]
		dBPerfInfoRepository.Queries.Add(query, extractedLog)
		return nil
	}
	return errors.New("error parsing message as LogMessageWithNewQueryRegEx")
}

func getNewQuery(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) stucts.PartialQuery {
	dBPerfInfoRepository.CurrentQuery = &stucts.Query{QueryId: extractedLog.QueryId}
	processIDs := types.InitIntSet()
	processIDs.Add(extractedLog.ThreadId)
	completed := true
	query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId, Timestamp: &extractedLog.Timestamp, ProcessIDs: &processIDs, Completed: &completed}
	return query
}

func ParseMessageWithQueryDurationV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithDurationRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 3 {
		if dBPerfInfoRepository.CurrentQuery == nil {
			query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId}
			query.Duration = parseDuration(parts[1], parts[2])
			dBPerfInfoRepository.Queries.Add(query, extractedLog)
			return nil
		}
	}
	return errors.New("error parsing message as LogMessageWithDurationRegEx")
}

func ParseMessageDurationAndQueryV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithDurationAndQueryRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 4 {
		if dBPerfInfoRepository.CurrentQuery == nil {
			dBPerfInfoRepository.CurrentQuery = &stucts.Query{QueryId: extractedLog.QueryId}
		}
		query := getNewQuery(extractedLog, dBPerfInfoRepository)
		query.Duration = parseDuration(parts[1], parts[2])
		query.Query = &parts[3]
		dBPerfInfoRepository.Queries.Add(query, extractedLog)
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

func ParseLogMessageWithSystemUsageRegExV1(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	if parts := PostgresLogMessageWithSystemUsageRegEx.FindStringSubmatch(extractedLog.Message); len(parts) == 3 {
		if dBPerfInfoRepository.CurrentQuery != nil {
			memoryUsage, _ := formattedSizeToBytes(parts[1], parts[2])
			query := stucts.PartialQuery{QueryId: &dBPerfInfoRepository.CurrentQuery.QueryId, PeakMemoryUsage: &memoryUsage}
			dBPerfInfoRepository.Queries.Add(query, extractedLog)
		}
		return nil
	}
	return errors.New("error parsing message as PostgresLogMessageWithSystemUsageRegEx")
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
	sizeInInt, _ := strconv.ParseInt(size, 0, 0)
	sizeInFloat := float64(sizeInInt)
	if unit == "B" {
		return sizeInFloat, nil
	} else if unit == "kB" {
		return sizeInFloat * 1024, nil
	} else if unit == "mB" {
		return sizeInFloat * 1024 * 1024, nil
	} else if unit == "gB" {
		return sizeInFloat * 1024 * 1024 * 1024, nil
	}
	return 0, errors.New("size format " + unit + " not supported")
}
