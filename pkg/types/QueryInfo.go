package types

import (
	"github.com/montanaflynn/stats"
	"sort"
	"time"
)

type QueryInfo struct {
	Query            string
	Count            int
	Durations        []float64
	ReadRows         []int
	ReadBytes        []float64
	Databases        []string
	Tables           []string
	Completed        []bool
	ClientHosts      []string
	Users            []string
	PeakMemoryUsages []float64
	Timestamps       []time.Time
}

func InitQueryInfo(queryString string) *QueryInfo {
	return &QueryInfo{

		Query:            queryString,
		Durations:        []float64{},
		ReadRows:         []int{},
		ReadBytes:        []float64{},
		Databases:        []string{},
		Tables:           []string{},
		Completed:        []bool{},
		ClientHosts:      []string{},
		Users:            []string{},
		PeakMemoryUsages: []float64{},
		Timestamps:       []time.Time{},
	}
}

func (singleQueryInfo *QueryInfo) Add(query *Query) {
	singleQueryInfo.Count += 1
	singleQueryInfo.Durations = append(singleQueryInfo.Durations, query.Duration)
	singleQueryInfo.ReadRows = append(singleQueryInfo.ReadRows, query.ReadRows)
	singleQueryInfo.ReadBytes = append(singleQueryInfo.ReadBytes, query.ReadBytes)
	//TODO check this logic, only one table and database is used
	databases := query.Databases
	if len(databases) > 1 {
		delete(databases, "*")
	}
	for _, database := range databases.ToArray() {
		singleQueryInfo.Databases = append(singleQueryInfo.Databases, database)
		break
	}
	tables := query.Tables
	if len(tables) > 1 {
		delete(tables, "*")
	}
	for _, table := range tables.ToArray() {
		singleQueryInfo.Tables = append(singleQueryInfo.Tables, table)
		break
	}
	singleQueryInfo.Completed = append(singleQueryInfo.Completed, query.Completed)
	singleQueryInfo.ClientHosts = append(singleQueryInfo.ClientHosts, query.ClientHost)
	singleQueryInfo.Users = append(singleQueryInfo.Users, query.User)
	singleQueryInfo.PeakMemoryUsages = append(singleQueryInfo.PeakMemoryUsages, query.PeakMemoryUsage)
	singleQueryInfo.Timestamps = append(singleQueryInfo.Timestamps, query.Timestamp)
}

func (singleQueryInfo *QueryInfo) CompleteProcessing() {
	sort.Float64s(singleQueryInfo.Durations)
	sort.Float64s(singleQueryInfo.PeakMemoryUsages)
	sort.Ints(singleQueryInfo.ReadRows)
	sort.Float64s(singleQueryInfo.ReadBytes)
	sort.Slice(singleQueryInfo.Timestamps, func(i, j int) bool {
		return singleQueryInfo.Timestamps[i].Before(singleQueryInfo.Timestamps[j])
	})
}

func (singleQueryInfo *QueryInfo) GetQPS() float64 {
	diff := singleQueryInfo.Timestamps[len(singleQueryInfo.Timestamps)-1].Sub(singleQueryInfo.Timestamps[0])
	if diff == 0 {
		return float64(singleQueryInfo.Count)
	}
	return float64(singleQueryInfo.Count) / diff.Seconds()
}

func (singleQueryInfo *QueryInfo) GetTotalDuration() float64 {
	var totalDuration float64
	for _, duration := range singleQueryInfo.Durations {
		totalDuration += duration
	}
	return totalDuration
}

func (singleQueryInfo *QueryInfo) GetMinDuration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		return singleQueryInfo.Durations[0]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMaxDuration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		return singleQueryInfo.Durations[len(singleQueryInfo.Durations)-1]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetAvgDuration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		return singleQueryInfo.GetTotalDuration() / float64(len(singleQueryInfo.Durations))
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetPercentile95Duration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		data := stats.LoadRawData(singleQueryInfo.Durations)
		percentile, _ := data.Percentile(95)
		return percentile
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMedianDuration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		data := stats.LoadRawData(singleQueryInfo.Durations)
		median, _ := data.Median()
		return median
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetStdDevDuration() float64 {
	if len(singleQueryInfo.Durations) > 0 {
		data := stats.LoadRawData(singleQueryInfo.Durations)
		stdDev, _ := data.StandardDeviation()
		return stdDev
	}
	return 0
}

// ReadRows
func (singleQueryInfo *QueryInfo) GetTotalReadRows() float64 {
	var total float64
	for _, i := range singleQueryInfo.ReadRows {
		total += float64(i)
	}
	return total
}

func (singleQueryInfo *QueryInfo) GetMinReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		return float64(singleQueryInfo.ReadRows[0])
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMaxReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		return float64(singleQueryInfo.ReadRows[len(singleQueryInfo.ReadRows)-1])
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetAvgReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		return singleQueryInfo.GetTotalReadRows() / float64(len(singleQueryInfo.ReadRows))
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetPercentile95ReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadRows)
		percentile, _ := data.Percentile(95)
		return percentile
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMedianReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadRows)
		median, _ := data.Median()
		return median
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetStdDevReadRows() float64 {
	if len(singleQueryInfo.ReadRows) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadRows)
		stdDev, _ := data.StandardDeviation()
		return stdDev
	}
	return 0
}

// ReadBytes
func (singleQueryInfo *QueryInfo) GetTotalReadBytes() float64 {
	var total float64
	for _, i := range singleQueryInfo.ReadBytes {
		total += i
	}
	return total
}

func (singleQueryInfo *QueryInfo) GetMinReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		return singleQueryInfo.ReadBytes[0]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMaxReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		return singleQueryInfo.ReadBytes[len(singleQueryInfo.ReadBytes)-1]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetAvgReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		return singleQueryInfo.GetTotalReadBytes() / float64(len(singleQueryInfo.ReadBytes))
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetPercentile95ReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadBytes)
		percentile, _ := data.Percentile(95)
		return percentile
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMedianReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadBytes)
		median, _ := data.Median()
		return median
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetStdDevReadBytes() float64 {
	if len(singleQueryInfo.ReadBytes) > 0 {
		data := stats.LoadRawData(singleQueryInfo.ReadBytes)
		stdDev, _ := data.StandardDeviation()
		return stdDev
	}
	return 0
}

// PeakMemoryUsage
func (singleQueryInfo *QueryInfo) GetTotalPeakMemoryUsage() float64 {
	var total float64
	for _, i := range singleQueryInfo.PeakMemoryUsages {
		total += i
	}
	return total
}

func (singleQueryInfo *QueryInfo) GetMinPeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		return singleQueryInfo.PeakMemoryUsages[0]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMaxPeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		return singleQueryInfo.PeakMemoryUsages[len(singleQueryInfo.PeakMemoryUsages)-1]
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetAvgPeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		return singleQueryInfo.GetTotalPeakMemoryUsage() / float64(len(singleQueryInfo.PeakMemoryUsages))
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetPercentile95PeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		data := stats.LoadRawData(singleQueryInfo.PeakMemoryUsages)
		percentile, _ := data.Percentile(95)
		return percentile
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetMedianPeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		data := stats.LoadRawData(singleQueryInfo.PeakMemoryUsages)
		median, _ := data.Median()
		return median
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetStdDevPeakMemoryUsage() float64 {
	if len(singleQueryInfo.PeakMemoryUsages) > 0 {
		data := stats.LoadRawData(singleQueryInfo.PeakMemoryUsages)
		stdDev, _ := data.StandardDeviation()
		return stdDev
	}
	return 0
}

func (singleQueryInfo *QueryInfo) GetHostsWithCount() StringCountPairArray {
	return getStringCountPair(singleQueryInfo.ClientHosts)
}

func (singleQueryInfo *QueryInfo) GetDatabasesWithCount() StringCountPairArray {
	return getStringCountPair(singleQueryInfo.Databases)
}

func (singleQueryInfo *QueryInfo) GetTablesWithCount() StringCountPairArray {
	return getStringCountPair(singleQueryInfo.Tables)
}

func (singleQueryInfo *QueryInfo) GetUsersWithCount() StringCountPairArray {
	return getStringCountPair(singleQueryInfo.Users)
}

func getStringCountPair(sa []string) StringCountPairArray {
	stringCountMap := map[string]int{}
	var stringCountPairArray StringCountPairArray
	for _, s := range sa {
		stringCountMap[s] += 1
	}
	for s, count := range stringCountMap {
		stringCountPairArray = append(stringCountPairArray, StringCountPair{
			String: s,
			Count:  count,
		})
	}

	return stringCountPairArray
}
