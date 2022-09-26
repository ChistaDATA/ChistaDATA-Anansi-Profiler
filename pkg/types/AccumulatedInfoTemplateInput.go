package types

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	"github.com/montanaflynn/stats"
	"os"
	"sort"
	"time"
)

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

func InitAccumulatedInfoTemplateInput(queryInfos []*QueryInfo, filePaths []string) AccumulatedInfoTemplateInput {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	var durations []float64
	var readRows []int
	var readBytes []float64
	var peakMemoryUsages []float64
	var timeStamps []time.Time
	totalQueryCount := 0

	for _, info := range queryInfos {
		totalQueryCount += info.Count
		durations = append(durations, info.Durations...)
		readRows = append(readRows, info.ReadRows...)
		readBytes = append(readBytes, info.ReadBytes...)
		peakMemoryUsages = append(peakMemoryUsages, info.PeakMemoryUsages...)
		timeStamps = append(timeStamps, info.Timestamps...)
	}

	sort.Slice(timeStamps, func(i, j int) bool {
		return timeStamps[i].Before(timeStamps[j])
	})

	return AccumulatedInfoTemplateInput{
		TotalDuration:         getAccumulatedTotalDuration(durations),
		CurrentDate:           time.Now().String(),
		Hostname:              hostname,
		Files:                 filePathsToString(filePaths),
		Duration:              getDurationInfo(durations),
		ReadRows:              getReadRowsInfo(readRows),
		ReadBytes:             getReadBytesInfo(readBytes),
		PeakMemoryUsage:       getPeakMemoryUsageInfo(peakMemoryUsages),
		FromTimestamp:         timeStamps[0].String(),
		ToTimestamp:           timeStamps[len(timeStamps)-1].String(),
		TotalQueryCount:       formatters.IntToKMilBilTri(totalQueryCount, 7),
		TotalUniqueQueryCount: formatters.IntToKMilBilTri(len(queryInfos), 7),
		TotalQPS:              getTotalQPS(totalQueryCount, timeStamps),
	}
}

func getTotalQPS(totalQueryCount int, timeStamps []time.Time) string {
	diff := timeStamps[len(timeStamps)-1].Sub(timeStamps[0])
	return formatters.Float64ToKMilBilTri(float64(totalQueryCount)/diff.Seconds(), 7)
}

func getAccumulatedTotalDuration(durations []float64) float64 {
	data := stats.LoadRawData(durations)
	sum, _ := data.Sum()
	return sum
}

func getPeakMemoryUsageInfo(dataArray []float64) QueryInfoTemplateInputPeakMemoryUsage {
	data := stats.LoadRawData(dataArray)
	sum, min, max, avg, per95, stdDev, median := FindMatrices(data)
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f, 7), 7)
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
	sum, min, max, avg, per95, stdDev, median := FindMatrices(data)
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ByteSizeToString(f, 7), 7)
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
	sum, min, max, avg, per95, stdDev, median := FindMatrices(data)
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64ToKMilBilTri(f, 7), 7)
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
	sum, min, max, avg, per95, stdDev, median := FindMatrices(data)
	formatToString := func(f float64) string {
		return formatters.PrefixSpace(formatters.Float64SecondsToString(f, 7), 7)
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

func FindMatrices(data stats.Float64Data) (float64, float64, float64, float64, float64, float64, float64) {
	sum, _ := data.Sum()
	min, _ := data.Min()
	max, _ := data.Max()
	avg, _ := data.Mean()
	per95, _ := data.Percentile(95)
	stdDev, _ := data.StandardDeviation()
	median, _ := data.Median()
	return sum, min, max, avg, per95, stdDev, median
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
