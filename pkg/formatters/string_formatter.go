package formatters

import (
	"fmt"
	"math"
	"sort"
)

//This file contains converters which will convert values to readable formats

var byteSizeExtensions = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

var secondsExtension = [...]string{"", "m", "µ", "n", "p"}

var siMultiplierPrefixes = map[float64]string{24: "Y", 21: "Z", 18: "E", 15: "P", 12: "T", 9: "G", 6: "M", 3: "k", 0: "", -3: "m", -6: "µ", -9: "n", -12: "p", -15: "f", -18: "a", -21: "z", -24: "y"}
var siMultiplierPrefixKeys []float64

var maxSiMultiplierPrefixVal float64
var minSiMultiplierPrefixVal float64

func init() {
	for k := range siMultiplierPrefixes {
		siMultiplierPrefixKeys = append(siMultiplierPrefixKeys, k)
		if k < minSiMultiplierPrefixVal {
			minSiMultiplierPrefixVal = k
		}

		if k > maxSiMultiplierPrefixVal {
			maxSiMultiplierPrefixVal = k
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(siMultiplierPrefixKeys)))
}

// Float64SecondsToString Converts Seconds to readable string format.
// s, ms, µs, ns, ps are possible values
func Float64SecondsToString(f float64) string {
	sf, u := Float64ToString(f)
	return fmt.Sprintf("%.2f%ss", sf, u)
}

// Float64ToString Converts Float to readable string format.
func Float64ToString(s float64) (float64, string) {
	prefixPower := 0.0
	if s != 0 && s < 0.001 {
		for s < 0.002 && prefixPower > minSiMultiplierPrefixVal {
			s *= 10
			prefixPower -= 1
		}
	} else {
		for s > 999 && prefixPower < maxSiMultiplierPrefixVal {
			s /= 10
			prefixPower += 1
		}
	}

	for i := 0; i < len(siMultiplierPrefixKeys); i++ {
		if prefixPower == siMultiplierPrefixKeys[i] {
			break
		} else if prefixPower > siMultiplierPrefixKeys[i] {
			if i > 0 {
				diffWithPrev := math.Abs(siMultiplierPrefixKeys[i-1] - prefixPower)
				for j := 0.0; j < diffWithPrev; j++ {
					s /= 10
				}
				prefixPower = siMultiplierPrefixKeys[i-1]
			}
			//diffWithCurrent := math.Abs(siMultiplierPrefixKeys[i] - prefixPower)
			//else {
			//for j := 0.0; j < diffWithCurrent; j++ {
			//	s /= 10
			//}
			//prefixPower = siMultiplierPrefixKeys[i]
			//}
			break
		}
	}

	return s, siMultiplierPrefixes[prefixPower]
}

// Float64ToNumberWithSIMultipliers Converts numbers to readable string format.
// eg: 54297892 will be converted to 54.29M
// K, M, B, T are possible values ( k->Thousands, M->Millions, B->Billions, T->Trillions)
func Float64ToNumberWithSIMultipliers(size float64) string {
	s, p := Float64ToString(size)
	return fmt.Sprintf("%.2f%s", s, p)
}

func IntToNumberWithSIMultipliers(size int) string {
	return Float64ToNumberWithSIMultipliers(float64(size))
}

// Float64ByteSizeToString Converts bytes to readable string format.
// eg: 54297892 bytes will be converted to 54.29MB
// B, KB, MB, GB, TB, PB, EB, ZB are possible values
func Float64ByteSizeToString(size float64) string {
	extensionIndex := 0
	for extensionIndex < len(byteSizeExtensions) && size >= 100 {
		size /= 1024
		extensionIndex += 1
	}
	if extensionIndex >= len(byteSizeExtensions) {
		extensionIndex = len(byteSizeExtensions) - 1
		return fmt.Sprintf("%.0f%s", size, byteSizeExtensions[extensionIndex])
	}

	return fmt.Sprintf("%.2f%s", size, byteSizeExtensions[extensionIndex])
}

func PrefixSpace(s string, limit int) string {
	for len(s) < limit {
		s = " " + s
	}
	return s[:limit]
}

// PercentageToCharRep function represent a certain percentage in a specific character as bar graph
func PercentageToCharRep(rep string, count int, total int, limit int) string {
	number := int((float64(count) / float64(total)) * float64(limit))
	s := ""
	for len(s) < number {
		s += rep
	}
	return s
}
