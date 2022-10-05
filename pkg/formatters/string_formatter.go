package formatters

import "fmt"

//This file contains converters which will convert values to readable formats

var byteSizeExtensions = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

var kMilBilTriExtensions = [...]string{"", "K", "M", "B", "T"}

var secondsExtension = [...]string{"", "m", "µ", "n", "p"}

// Float64SecondsToString Converts Seconds to readable string format.
// s, ms, µs, ns, ps are possible values
func Float64SecondsToString(s float64) string {
	extensionIndex := 0
	for s < 0 && extensionIndex < len(secondsExtension) {
		s *= 1000
		extensionIndex += 1
	}
	if extensionIndex >= len(secondsExtension) {
		extensionIndex = len(secondsExtension) - 1
		return fmt.Sprintf("%.2f%ss", s, secondsExtension[extensionIndex])
	}
	return fmt.Sprintf("%.2f%ss", s, secondsExtension[extensionIndex])
}

// Float64ToKMilBilTri Converts numbers to readable string format.
// eg: 54297892 will be converted to 54.29M
// K, M, B, T are possible values ( k->Thousands, M->Millions, B->Billions, T->Trillions)
func Float64ToKMilBilTri(size float64) string {
	extensionIndex := 0
	for extensionIndex < len(kMilBilTriExtensions) && size >= 100 {
		size /= 1000
		extensionIndex += 1
	}
	if extensionIndex >= len(kMilBilTriExtensions) {
		extensionIndex = len(kMilBilTriExtensions) - 1
		return fmt.Sprintf("%.0f%s", size, kMilBilTriExtensions[extensionIndex])
	}
	return fmt.Sprintf("%.2f%s", size, kMilBilTriExtensions[extensionIndex])
}

func IntToKMilBilTri(size int) string {
	return Float64ToKMilBilTri(float64(size))
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
