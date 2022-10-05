package utils

import (
	"github.com/montanaflynn/stats"
)

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
