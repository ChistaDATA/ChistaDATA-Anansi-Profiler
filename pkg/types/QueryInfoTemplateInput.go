package types

import (
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	"sort"
	"strconv"
)

type QueryInfoTemplateInputQueryTimeDistribution struct {
	Under10us          string
	Over10usUnder100us string
	Over100usUnder1ms  string
	Over1msUnder10ms   string
	Over10msUnder100ms string
	Over100msUnder1s   string
	Over1sUnder10s     string
	Over10s            string
}

func InitQueryInfoTemplateInputQueryTimeDistribution(info *QueryInfo) QueryInfoTemplateInputQueryTimeDistribution {
	formatToString := func(c int) string {
		return formatters.PercentageToCharRep("#", c, info.Count, 60)
	}
	var counts [8]int
	di := 0
	ci := 0
	for di < len(info.Durations) && info.Durations[di] < 0.00001 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 0.0001 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 0.001 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 0.01 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 0.1 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 1 {
		counts[ci]++
		di++
	}
	ci++
	for di < len(info.Durations) && info.Durations[di] < 10 {
		counts[ci]++
		di++
	}
	ci++
	counts[ci] = len(info.Durations) - di
	return QueryInfoTemplateInputQueryTimeDistribution{
		Under10us:          formatToString(counts[0]),
		Over10usUnder100us: formatToString(counts[1]),
		Over100usUnder1ms:  formatToString(counts[2]),
		Over1msUnder10ms:   formatToString(counts[3]),
		Over10msUnder100ms: formatToString(counts[4]),
		Over100msUnder1s:   formatToString(counts[5]),
		Over1sUnder10s:     formatToString(counts[6]),
		Over10s:            formatToString(counts[7]),
	}
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

func InitQueryInfoTemplateInputPeakMemoryUsage(info *QueryInfo) QueryInfoTemplateInputPeakMemoryUsage {
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f, 7), 7)
	}
	return QueryInfoTemplateInputPeakMemoryUsage{
		Total:        formatToString(info.GetTotalPeakMemoryUsage()),
		Min:          formatToString(info.GetMinPeakMemoryUsage()),
		Max:          formatToString(info.GetMaxPeakMemoryUsage()),
		Avg:          formatToString(info.GetAvgPeakMemoryUsage()),
		Percentile95: formatToString(info.GetPercentile95PeakMemoryUsage()),
		StdDev:       formatToString(info.GetStdDevPeakMemoryUsage()),
		Median:       formatToString(info.GetMedianPeakMemoryUsage()),
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

func InitQueryInfoTemplateInputReadBytes(info *QueryInfo) QueryInfoTemplateInputReadBytes {
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f, 7), 7)
	}
	return QueryInfoTemplateInputReadBytes{
		Total:        formatToString(info.GetTotalReadBytes()),
		Min:          formatToString(info.GetMinReadBytes()),
		Max:          formatToString(info.GetMaxReadBytes()),
		Avg:          formatToString(info.GetAvgReadBytes()),
		Percentile95: formatToString(info.GetPercentile95ReadBytes()),
		StdDev:       formatToString(info.GetStdDevReadBytes()),
		Median:       formatToString(info.GetMedianReadBytes()),
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

func InitQueryInfoTemplateInputReadRows(info *QueryInfo) QueryInfoTemplateInputReadRows {
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ToKMilBilTri(f, 7), 7)
	}
	return QueryInfoTemplateInputReadRows{
		Total:        formatToString(info.GetTotalReadRows()),
		Min:          formatToString(info.GetMinReadRows()),
		Max:          formatToString(info.GetMaxReadRows()),
		Avg:          formatToString(info.GetAvgReadRows()),
		Percentile95: formatToString(info.GetPercentile95ReadRows()),
		StdDev:       formatToString(info.GetStdDevReadRows()),
		Median:       formatToString(info.GetMedianReadRows()),
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

func InitQueryInfoTemplateInputDuration(info *QueryInfo) QueryInfoTemplateInputDuration {
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64SecondsToString(f, 7), 7)
	}
	return QueryInfoTemplateInputDuration{
		Total:        formatToString(info.GetTotalDuration()),
		Min:          formatToString(info.GetMinDuration()),
		Max:          formatToString(info.GetMaxDuration()),
		Avg:          formatToString(info.GetAvgDuration()),
		Percentile95: formatToString(info.GetPercentile95Duration()),
		StdDev:       formatToString(info.GetStdDevDuration()),
		Median:       formatToString(info.GetMedianDuration()),
	}
}

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
}

func InitQueryInfoTemplateInput(index int, info *QueryInfo) QueryInfoTemplateInput {
	return QueryInfoTemplateInput{
		Query:                 info.Query,
		Pos:                   strconv.Itoa(index + 1),
		QPS:                   formatters.Float64ToKMilBilTri(info.GetQPS(), 7),
		FromTimestamp:         info.Timestamps[0].String(),
		ToTimestamp:           info.Timestamps[len(info.Timestamps)-1].String(),
		Count:                 formatters.PrefixSpace(formatters.IntToKMilBilTri(info.Count, 7), 7),
		Duration:              InitQueryInfoTemplateInputDuration(info),
		ReadRows:              InitQueryInfoTemplateInputReadRows(info),
		ReadBytes:             InitQueryInfoTemplateInputReadBytes(info),
		PeakMemoryUsage:       InitQueryInfoTemplateInputPeakMemoryUsage(info),
		QueryTimeDistribution: InitQueryInfoTemplateInputQueryTimeDistribution(info),
		HostInfo:              getHostInfo(info),
		DatabaseInfo:          getDatabaseInfo(info),
		TableInfo:             getTableInfo(info),
		UserInfo:              getUserInfo(info),
	}
}

func getHostInfo(info *QueryInfo) string {
	hostWithCount := info.GetHostsWithCount()
	sort.Sort(hostWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, hostWithCount, limit)
}

func getDatabaseInfo(info *QueryInfo) string {
	databaseWithCount := info.GetDatabasesWithCount()
	sort.Sort(databaseWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, databaseWithCount, limit)
}

func getTableInfo(info *QueryInfo) string {
	tableWithCount := info.GetDatabasesWithCount()
	sort.Sort(tableWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, tableWithCount, limit)
}

func getUserInfo(info *QueryInfo) string {
	usersWithCount := info.GetUsersWithCount()
	sort.Sort(usersWithCount)
	limit := 3
	return formatStringCountPairArrayWithLimit(info, usersWithCount, limit)
}

func formatStringCountPairArrayWithLimit(info *QueryInfo, array StringCountPairArray, limit int) string {
	if limit >= len(array) {
		limit = len(array)
	}

	arrayString := ""
	for _, pair := range array {
		arrayString += fmt.Sprintf("%s (%d/%d)  ", pair.String, pair.Count, info.Count)
	}
	return arrayString
}
