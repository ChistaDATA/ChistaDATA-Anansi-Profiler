package stucts

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/utils"
	"github.com/montanaflynn/stats"
	"sort"
	"time"
)

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
	if similarQueryInfo.FromTimestamp == nil || query.Timestamp.Before(*similarQueryInfo.FromTimestamp) {
		t := query.Timestamp
		similarQueryInfo.FromTimestamp = &t
	}
	if similarQueryInfo.ToTimestamp == nil || query.Timestamp.After(*similarQueryInfo.ToTimestamp) {
		t := query.Timestamp
		similarQueryInfo.ToTimestamp = &t
	}
}

func (similarQueryInfo *SimilarQueryInfo) CompleteProcessing() {
	sort.Float64s(similarQueryInfo.Durations)
	sort.Float64s(similarQueryInfo.PeakMemoryUsages)
	sort.Ints(similarQueryInfo.ReadRows)
	sort.Float64s(similarQueryInfo.ReadBytes)
}

func (similarQueryInfo *SimilarQueryInfo) GetQPS(totalDuration float64) float64 {
	diff := similarQueryInfo.ToTimestamp.Sub(*similarQueryInfo.FromTimestamp)
	if diff.Seconds() == 0 {
		return float64(similarQueryInfo.Count)
	}
	return float64(similarQueryInfo.Count) / totalDuration
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

func (similarQueryInfo *SimilarQueryInfo) GetMaxDuration() float64 {
	var max float64
	for i := 0; i < len(similarQueryInfo.Durations); i++ {
		if max < similarQueryInfo.Durations[i] {
			max = similarQueryInfo.Durations[i]
		}
	}
	return max
}

func (similarQueryInfo *SimilarQueryInfo) GetTotalDuration() float64 {
	var sum float64
	for i := 0; i < len(similarQueryInfo.Durations); i++ {
		sum += similarQueryInfo.Durations[i]
	}
	return sum
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
