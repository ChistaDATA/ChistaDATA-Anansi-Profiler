package stucts

import (
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/utils"
	"github.com/montanaflynn/stats"
	"os"
	"sort"
	"time"
)

// AccumulatedInfoTemplateInput Input struct for AccumulatedInfoTemplate
type AccumulatedInfoTemplateInput struct {
	CurrentDate           string
	Hostname              string
	Files                 string
	TotalQueryCount       string
	TotalUniqueQueryCount string
	TotalQPS              string
	FromTimestamp         string
	ToTimestamp           string
	Duration              QueryInfoTemplateInputDuration
	ReadRows              QueryInfoTemplateInputReadRows
	ReadBytes             QueryInfoTemplateInputReadBytes
	PeakMemoryUsage       QueryInfoTemplateInputPeakMemoryUsage
	TotalDuration         float64
}

func InitAccumulatedInfoTemplateInput(queryInfos SimilarQueryInfoList, filePaths []string) AccumulatedInfoTemplateInput {
	hostname, _ := os.Hostname()

	var durations []float64
	var readRows []int
	var readBytes []float64
	var peakMemoryUsages []float64
	var timeStamps []*time.Time
	totalQueryCount := 0
	t := time.Time{}

	for _, info := range queryInfos {
		totalQueryCount += info.Count
		durations = append(durations, info.Durations...)
		readRows = append(readRows, info.ReadRows...)
		readBytes = append(readBytes, info.ReadBytes...)
		peakMemoryUsages = append(peakMemoryUsages, info.PeakMemoryUsages...)
		if *info.FromTimestamp != t {
			timeStamps = append(timeStamps, info.FromTimestamp)
		}

		if *info.ToTimestamp != t {
			timeStamps = append(timeStamps, info.ToTimestamp)
		}
	}

	sort.Slice(timeStamps, func(i, j int) bool {
		return (*timeStamps[i]).Before(*timeStamps[j])
	})

	for len(timeStamps) < 2 {
		timeStamps = append(timeStamps, &t)
	}

	return AccumulatedInfoTemplateInput{
		//TotalDuration:         getAccumulatedTotalDuration(durations),
		TotalDuration:         (*timeStamps[len(timeStamps)-1]).Sub(*timeStamps[0]).Seconds(),
		CurrentDate:           time.Now().String(),
		Hostname:              hostname,
		Files:                 filePathsToString(filePaths),
		Duration:              getDurationInfo(durations),
		ReadRows:              getReadRowsInfo(readRows),
		ReadBytes:             getReadBytesInfo(readBytes),
		PeakMemoryUsage:       getPeakMemoryUsageInfo(peakMemoryUsages),
		FromTimestamp:         timeStamps[0].String(),
		ToTimestamp:           timeStamps[len(timeStamps)-1].String(),
		TotalQueryCount:       fmt.Sprintf("%d", totalQueryCount),
		TotalUniqueQueryCount: fmt.Sprintf("%d", len(queryInfos)),
		TotalQPS:              getTotalQPS(totalQueryCount, timeStamps),
	}
}

func getTotalQPS(totalQueryCount int, timeStamps []*time.Time) string {
	diff := (*timeStamps[len(timeStamps)-1]).Sub(*timeStamps[0])
	if diff == 0 {
		formatters.Float64ToNumberWithSIMultipliers(float64(totalQueryCount))
	}
	return formatters.Float64ToNumberWithSIMultipliers(float64(totalQueryCount) / diff.Seconds())
}

func getAccumulatedTotalDuration(durations []float64) float64 {
	data := stats.LoadRawData(durations)
	sum, _ := data.Sum()
	return sum
}

func getPeakMemoryUsageInfo(dataArray []float64) QueryInfoTemplateInputPeakMemoryUsage {
	data := stats.LoadRawData(dataArray)
	sum, min, max, avg, per95, stdDev, median := utils.FindMatrices(data)
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

func getReadBytesInfo(dataArray []float64) QueryInfoTemplateInputReadBytes {
	data := stats.LoadRawData(dataArray)
	sum, min, max, avg, per95, stdDev, median := utils.FindMatrices(data)
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

func getReadRowsInfo(dataArray []int) QueryInfoTemplateInputReadRows {
	data := stats.LoadRawData(dataArray)
	sum, min, max, avg, per95, stdDev, median := utils.FindMatrices(data)
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

func getDurationInfo(dataArray []float64) QueryInfoTemplateInputDuration {
	data := stats.LoadRawData(dataArray)
	sum, min, max, avg, per95, stdDev, median := utils.FindMatrices(data)
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

func filePathsToString(filePaths []string) string {
	filePathsString := ""
	if len(filePaths) > 0 {
		filePathsString = "\t* " + filePaths[0]
		for i := 1; i < len(filePaths); i++ {
			filePathsString = filePathsString + "\n" + "\t* " + filePaths[i]
		}
	}
	return filePathsString
}
