package stucts

import (
	"sort"
	"time"

	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/utils"
	"github.com/montanaflynn/stats"
)

// SimilarQueryInfo all important information of similar queries
type SimilarQueryInfo struct {
	Query            string
	Count            int
	Durations        []float64
	ReadRows         []int
	ReadBytes        []float64
	Databases        []string
	Tables           []string
	Completed        int
	ClientHosts      []string
	Users            []string
	PeakMemoryUsages []float64
	FromTimestamp    *time.Time
	ToTimestamp      *time.Time
	ErrorMessages    []string
}

func InitSimilarQueryInfo(queryString string) *SimilarQueryInfo {
	return &SimilarQueryInfo{
		Query:            queryString,
		Durations:        []float64{},
		ReadRows:         []int{},
		ReadBytes:        []float64{},
		Databases:        []string{},
		Tables:           []string{},
		ClientHosts:      []string{},
		Users:            []string{},
		PeakMemoryUsages: []float64{},
		FromTimestamp:    nil,
		ToTimestamp:      nil,
		ErrorMessages:    []string{},
	}
}

func (similarQueryInfo *SimilarQueryInfo) Add(query *Query) {
	similarQueryInfo.Count += 1
	similarQueryInfo.Durations = append(similarQueryInfo.Durations, query.Duration)
	similarQueryInfo.ReadRows = append(similarQueryInfo.ReadRows, query.ReadRows)
	similarQueryInfo.ReadBytes = append(similarQueryInfo.ReadBytes, query.ReadBytes)
	//TODO check this logic, only one table and database is used
	databases := query.Databases
	if len(databases) > 1 {
		delete(databases, "*")
	}
	for _, database := range databases.ToArray() {
		similarQueryInfo.Databases = append(similarQueryInfo.Databases, database)
		break
	}
	tables := query.Tables
	if len(tables) > 1 {
		delete(tables, "*")
	}
	for _, table := range tables.ToArray() {
		similarQueryInfo.Tables = append(similarQueryInfo.Tables, table)
		break
	}
	if query.Completed {
		similarQueryInfo.Completed += 1
	}
	similarQueryInfo.ClientHosts = append(similarQueryInfo.ClientHosts, query.ClientHost)
	similarQueryInfo.Users = append(similarQueryInfo.Users, query.User)
	similarQueryInfo.PeakMemoryUsages = append(similarQueryInfo.PeakMemoryUsages, query.PeakMemoryUsage)
	if similarQueryInfo.FromTimestamp == nil || (query.Timestamp != time.Time{} && query.Timestamp.Before(*similarQueryInfo.FromTimestamp)) {
		t := query.Timestamp
		similarQueryInfo.FromTimestamp = &t
	}
	if similarQueryInfo.ToTimestamp == nil || (query.Timestamp != time.Time{} && query.Timestamp.After(*similarQueryInfo.ToTimestamp)) {
		t := query.Timestamp
		similarQueryInfo.ToTimestamp = &t
	}
	if query.ErrorCompleteText != "" {
		similarQueryInfo.ErrorMessages = append(similarQueryInfo.ErrorMessages, query.ErrorCompleteText)
	}
}

func (similarQueryInfo *SimilarQueryInfo) CompleteProcessing() {
	sort.Float64s(similarQueryInfo.Durations)
	sort.Float64s(similarQueryInfo.PeakMemoryUsages)
	sort.Ints(similarQueryInfo.ReadRows)
	sort.Float64s(similarQueryInfo.ReadBytes)
}

func GetQPS(similarQueryInfo *SimilarQueryInfo, totalDuration float64) float64 {
	//diff := similarQueryInfo.ToTimestamp.Sub(*similarQueryInfo.FromTimestamp)
	//if diff.Seconds() == 0 {
	//	return float64(similarQueryInfo.Count)
	//}
	return float64(similarQueryInfo.Count) / totalDuration
}

//func GetQPS(similarQueryInfo *SimilarQueryInfo) float64 {
//	diff := similarQueryInfo.ToTimestamp.Sub(*similarQueryInfo.FromTimestamp)
//	if diff.Seconds() == 0 {
//		return float64(similarQueryInfo.Count)
//	}
//	return float64(similarQueryInfo.Count) / GetTotalDuration(similarQueryInfo)
//}

func GetCount(similarQueryInfo *SimilarQueryInfo) float64 {
	return float64(similarQueryInfo.Count)
}

func (similarQueryInfo *SimilarQueryInfo) GetDurationMatrices() (float64, float64, float64, float64, float64, float64, float64) {
	return utils.FindMatrices(stats.LoadRawData(similarQueryInfo.Durations))
}

func (similarQueryInfo *SimilarQueryInfo) GetReadRowsMatrices() (float64, float64, float64, float64, float64, float64, float64) {
	return utils.FindMatrices(stats.LoadRawData(similarQueryInfo.ReadRows))
}

func (similarQueryInfo *SimilarQueryInfo) GetReadBytesMatrices() (float64, float64, float64, float64, float64, float64, float64) {
	return utils.FindMatrices(stats.LoadRawData(similarQueryInfo.ReadBytes))
}

func (similarQueryInfo *SimilarQueryInfo) GetPeakMemoryUsageMatrices() (float64, float64, float64, float64, float64, float64, float64) {
	return utils.FindMatrices(stats.LoadRawData(similarQueryInfo.PeakMemoryUsages))
}

func (similarQueryInfo *SimilarQueryInfo) GetHostsWithCount() types.StringCountPairArray {
	return getStringCountPair(similarQueryInfo.ClientHosts)
}

func (similarQueryInfo *SimilarQueryInfo) GetDatabasesWithCount() types.StringCountPairArray {
	return getStringCountPair(similarQueryInfo.Databases)
}

func (similarQueryInfo *SimilarQueryInfo) GetTablesWithCount() types.StringCountPairArray {
	return getStringCountPair(similarQueryInfo.Tables)
}

func (similarQueryInfo *SimilarQueryInfo) GetUsersWithCount() types.StringCountPairArray {
	return getStringCountPair(similarQueryInfo.Users)
}

func (similarQueryInfo *SimilarQueryInfo) GetCompletedCount() int {
	return similarQueryInfo.Completed
}

func GetTotalDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	sum, _ := stats.LoadRawData(similarQueryInfo.Durations).Sum()
	return sum
}

func GetMaxDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	max, _ := stats.LoadRawData(similarQueryInfo.Durations).Max()
	return max
}

func GetMinDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	min, _ := stats.LoadRawData(similarQueryInfo.Durations).Min()
	return min
}

func GetAvgDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	avg, _ := stats.LoadRawData(similarQueryInfo.Durations).Min()
	return avg
}

func GetPer95Duration(similarQueryInfo *SimilarQueryInfo) float64 {
	per95, _ := stats.LoadRawData(similarQueryInfo.Durations).Percentile(95)
	return per95
}

func GetStdDevDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	stdDev, _ := stats.LoadRawData(similarQueryInfo.Durations).StandardDeviation()
	return stdDev
}

func GetMedianDuration(similarQueryInfo *SimilarQueryInfo) float64 {
	median, _ := stats.LoadRawData(similarQueryInfo.Durations).Median()
	return median
}

func GetTotalReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	sum, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Sum()
	return sum
}

func GetMaxReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	max, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Max()
	return max
}

func GetMinReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	min, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Min()
	return min
}

func GetAvgReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	avg, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Min()
	return avg
}

func GetPer95ReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	per95, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Percentile(95)
	return per95
}

func GetStdDevReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	stdDev, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).StandardDeviation()
	return stdDev
}

func GetMedianReadBytes(similarQueryInfo *SimilarQueryInfo) float64 {
	median, _ := stats.LoadRawData(similarQueryInfo.ReadBytes).Median()
	return median
}

func GetTotalReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	sum, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Sum()
	return sum
}

func GetMaxReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	max, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Max()
	return max
}

func GetMinReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	min, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Min()
	return min
}

func GetAvgReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	avg, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Min()
	return avg
}

func GetPer95ReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	per95, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Percentile(95)
	return per95
}

func GetStdDevReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	stdDev, _ := stats.LoadRawData(similarQueryInfo.ReadRows).StandardDeviation()
	return stdDev
}

func GetMedianReadRows(similarQueryInfo *SimilarQueryInfo) float64 {
	median, _ := stats.LoadRawData(similarQueryInfo.ReadRows).Median()
	return median
}

func GetTotalPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	sum, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Sum()
	return sum
}

func GetMaxPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	max, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Max()
	return max
}

func GetMinPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	min, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Min()
	return min
}

func GetAvgPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	avg, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Min()
	return avg
}

func GetPer95PeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	per95, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Percentile(95)
	return per95
}

func GetStdDevPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	stdDev, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).StandardDeviation()
	return stdDev
}

func GetMedianPeakMemory(similarQueryInfo *SimilarQueryInfo) float64 {
	median, _ := stats.LoadRawData(similarQueryInfo.PeakMemoryUsages).Median()
	return median
}

func getStringCountPair(sa []string) types.StringCountPairArray {
	stringCountMap := map[string]int{}
	var stringCountPairArray types.StringCountPairArray
	for _, s := range sa {
		stringCountMap[s] += 1
	}
	for s, count := range stringCountMap {
		stringCountPairArray = append(stringCountPairArray, types.StringCountPair{
			String: s,
			Count:  count,
		})
	}

	return stringCountPairArray
}
