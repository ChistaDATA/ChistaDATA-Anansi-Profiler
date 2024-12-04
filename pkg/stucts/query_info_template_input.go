package stucts

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
)

// QueryInfoTemplateInput input for QueryInfoTemplate
type QueryInfoTemplateInput struct {
	Query                 string
	Pos                   string
	QPS                   string
	FromTimestamp         string
	ToTimestamp           string
	Count                 string
	Duration              QueryInfoTemplateInputDuration
	ReadRows              QueryInfoTemplateInputReadRows
	ReadBytes             QueryInfoTemplateInputReadBytes
	PeakMemoryUsage       QueryInfoTemplateInputPeakMemoryUsage
	QueryTimeDistribution QueryInfoTemplateInputQueryTimeDistribution
	HostInfo              string
	DatabaseInfo          string
	UserInfo              string
	TableInfo             string
	CompletedInfo         string
	ErrorInfo             string
}

func InitQueryInfoTemplateInput(index int, info *SimilarQueryInfo, totalDuration float64) QueryInfoTemplateInput {
	return QueryInfoTemplateInput{
		Query:                 info.Query,
		Pos:                   strconv.Itoa(index + 1),
		QPS:                   fmt.Sprintf("%.03f", GetQPS(info, totalDuration)),
		FromTimestamp:         info.FromTimestamp.String(),
		ToTimestamp:           info.ToTimestamp.String(),
		Count:                 formatters.PrefixSpace(formatters.IntToNumberWithSIMultipliers(info.Count), 7),
		Duration:              InitQueryInfoTemplateInputDuration(info),
		ReadRows:              InitQueryInfoTemplateInputReadRows(info),
		ReadBytes:             InitQueryInfoTemplateInputReadBytes(info),
		PeakMemoryUsage:       InitQueryInfoTemplateInputPeakMemoryUsage(info),
		QueryTimeDistribution: InitQueryInfoTemplateInputQueryTimeDistribution(info),
		HostInfo:              getHostInfo(info),
		DatabaseInfo:          getDatabaseInfo(info),
		TableInfo:             getTableInfo(info),
		UserInfo:              getUserInfo(info),
		CompletedInfo:         getCompletedInfo(info),
		ErrorInfo:             getErrorInfo(info),
	}
}

type QueryTimeDistributionString struct {
	Under10us          string
	Over10usUnder100us string
	Over100usUnder1ms  string
	Over1msUnder10ms   string
	Over10msUnder100ms string
	Over100msUnder1s   string
	Over1sUnder10s     string
	Over10s            string
}
type QueryTimeDistributionNumber struct {
	Under10us          int
	Over10usUnder100us int
	Over100usUnder1ms  int
	Over1msUnder10ms   int
	Over10msUnder100ms int
	Over100msUnder1s   int
	Over1sUnder10s     int
	Over10s            int
}

type QueryInfoTemplateInputQueryTimeDistribution struct {
	TimeDistString QueryTimeDistributionString
	TimeDistNumber QueryTimeDistributionNumber
}

func InitQueryInfoTemplateInputQueryTimeDistribution(info *SimilarQueryInfo) QueryInfoTemplateInputQueryTimeDistribution {
	formatToString := func(c int) string {
		return formatters.PercentageToCharRep("#", c, info.Count, 60)
	}
	var counts [8]int
	for _, duration := range info.Durations {
		if duration < 0.00001 {
			counts[0]++
		} else if duration < 0.0001 {
			counts[1]++
		} else if duration < 0.001 {
			counts[2]++
		} else if duration < 0.01 {
			counts[3]++
		} else if duration < 0.1 {
			counts[4]++
		} else if duration < 1 {
			counts[5]++
		} else if duration < 10 {
			counts[6]++
		} else {
			counts[7]++
		}
	}
	qts := QueryTimeDistributionString{
		Under10us:          formatToString(counts[0]),
		Over10usUnder100us: formatToString(counts[1]),
		Over100usUnder1ms:  formatToString(counts[2]),
		Over1msUnder10ms:   formatToString(counts[3]),
		Over10msUnder100ms: formatToString(counts[4]),
		Over100msUnder1s:   formatToString(counts[5]),
		Over1sUnder10s:     formatToString(counts[6]),
		Over10s:            formatToString(counts[7]),
	}
	qtn := QueryTimeDistributionNumber{
		Under10us:          counts[0],
		Over10usUnder100us: counts[1],
		Over100usUnder1ms:  counts[2],
		Over1msUnder10ms:   counts[3],
		Over10msUnder100ms: counts[4],
		Over100msUnder1s:   counts[5],
		Over1sUnder10s:     counts[6],
		Over10s:            counts[7],
	}
	return QueryInfoTemplateInputQueryTimeDistribution{qts, qtn}
}

type QueryInfoTemplateInputPeakMemoryUsage struct {
	Total        string
	Min          string
	Max          string
	Avg          string
	Percentile95 string
	StdDev       string
	Median       string
}

func InitQueryInfoTemplateInputPeakMemoryUsage(info *SimilarQueryInfo) QueryInfoTemplateInputPeakMemoryUsage {
	sum, min, max, avg, per95, stdDev, median := info.GetPeakMemoryUsageMatrices()
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f), 7)
	}
	return QueryInfoTemplateInputPeakMemoryUsage{
		Total:        formatToString(sum),
		Min:          formatToString(min),
		Max:          formatToString(max),
		Avg:          formatToString(avg),
		Percentile95: formatToString(per95),
		StdDev:       formatToString(stdDev),
		Median:       formatToString(median),
	}
}

type QueryInfoTemplateInputReadBytes struct {
	Total        string
	Min          string
	Max          string
	Avg          string
	Percentile95 string
	StdDev       string
	Median       string
}

func InitQueryInfoTemplateInputReadBytes(info *SimilarQueryInfo) QueryInfoTemplateInputReadBytes {
	sum, min, max, avg, per95, stdDev, median := info.GetReadBytesMatrices()
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f), 7)
	}
	return QueryInfoTemplateInputReadBytes{
		Total:        formatToString(sum),
		Min:          formatToString(min),
		Max:          formatToString(max),
		Avg:          formatToString(avg),
		Percentile95: formatToString(per95),
		StdDev:       formatToString(stdDev),
		Median:       formatToString(median),
	}
}

type QueryInfoTemplateInputReadRows struct {
	Total        string
	Min          string
	Max          string
	Avg          string
	Percentile95 string
	StdDev       string
	Median       string
}

func InitQueryInfoTemplateInputReadRows(info *SimilarQueryInfo) QueryInfoTemplateInputReadRows {
	sum, min, max, avg, per95, stdDev, median := info.GetReadRowsMatrices()
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ToNumberWithSIMultipliers(f), 7)
	}
	return QueryInfoTemplateInputReadRows{
		Total:        formatToString(sum),
		Min:          formatToString(min),
		Max:          formatToString(max),
		Avg:          formatToString(avg),
		Percentile95: formatToString(per95),
		StdDev:       formatToString(stdDev),
		Median:       formatToString(median),
	}
}

type QueryInfoTemplateInputDuration struct {
	Total        string
	Min          string
	Max          string
	Avg          string
	Percentile95 string
	StdDev       string
	Median       string
}

func InitQueryInfoTemplateInputDuration(info *SimilarQueryInfo) QueryInfoTemplateInputDuration {
	sum, min, max, avg, per95, stdDev, median := info.GetDurationMatrices()
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64SecondsToString(f), 7)
	}
	return QueryInfoTemplateInputDuration{
		Total:        formatToString(sum),
		Min:          formatToString(min),
		Max:          formatToString(max),
		Avg:          formatToString(avg),
		Percentile95: formatToString(per95),
		StdDev:       formatToString(stdDev),
		Median:       formatToString(median),
	}
}

func getCompletedInfo(info *SimilarQueryInfo) string {
	return fmt.Sprintf("%d/%d", info.GetCompletedCount(), info.Count)
}

func getHostInfo(info *SimilarQueryInfo) string {
	hostWithCount := info.GetHostsWithCount()
	sort.Sort(hostWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, hostWithCount, limit)
}

func getDatabaseInfo(info *SimilarQueryInfo) string {
	databaseWithCount := info.GetDatabasesWithCount()
	sort.Sort(databaseWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, databaseWithCount, limit)
}

func getTableInfo(info *SimilarQueryInfo) string {
	tableWithCount := info.GetDatabasesWithCount()
	sort.Sort(tableWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, tableWithCount, limit)
}

func getUserInfo(info *SimilarQueryInfo) string {
	usersWithCount := info.GetUsersWithCount()
	sort.Sort(usersWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, usersWithCount, limit)
}

func getErrorInfo(info *SimilarQueryInfo) string {

	return fmt.Sprintf("%d/%d", len(info.ErrorMessages), info.Count)
}

func formatStringCountPairArrayWithLimit(info *SimilarQueryInfo, array types.StringCountPairArray, limit int) string {
	if limit >= len(array) {
		limit = len(array)
	}

	arrayString := ""
	for _, pair := range array {
		arrayString += fmt.Sprintf("%s (%d/%d)  ", pair.String, pair.Count, info.Count)
	}
	return arrayString
}
