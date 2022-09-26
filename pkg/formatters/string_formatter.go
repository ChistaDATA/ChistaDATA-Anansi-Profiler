package formatters

import "fmt"

// float64Faction is 2 , +1 is for the . is 12.23
const float64Faction = 3

var byteSizeExtensions = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

var kMilBilTriExtensions = [...]string{"", "K", "M", "B", "T"}

var secondsExtension = [...]string{"", "m", "µ", "n", "p"}

func Float64SecondsToString(s float64, limit int) string {
	extensionIndex := 0
	for noOfDigits(int(s))+len(secondsExtension[extensionIndex])+float64Faction > limit && extensionIndex < len(secondsExtension) {
		s /= 1000
		extensionIndex += 1
	}
	if extensionIndex >= len(secondsExtension) {
		extensionIndex = len(secondsExtension) - 1
		return fmt.Sprintf("%.0f%ss", s, secondsExtension[extensionIndex])
	}
	return fmt.Sprintf("%.2f%ss", s, secondsExtension[extensionIndex])
}

func Float64ToKMilBilTri(size float64, limit int) string {
	extensionIndex := 0
	for noOfDigits(int(size))+len(kMilBilTriExtensions[extensionIndex])+float64Faction > limit && extensionIndex < len(kMilBilTriExtensions) {
		size /= 1000
		extensionIndex += 1
	}
	if extensionIndex >= len(kMilBilTriExtensions) {
		extensionIndex = len(kMilBilTriExtensions) - 1
		return fmt.Sprintf("%.0f%s", size, kMilBilTriExtensions[extensionIndex])
	}
	return fmt.Sprintf("%.2f%s", size, kMilBilTriExtensions[extensionIndex])
}

func IntToKMilBilTri(size int, limit int) string {
	extensionIndex := 0
	for noOfDigits(size)+len(kMilBilTriExtensions[extensionIndex]) > limit && extensionIndex < len(kMilBilTriExtensions) {
		size /= 1000
		extensionIndex += 1
	}
	if extensionIndex >= len(kMilBilTriExtensions) {
		extensionIndex = len(kMilBilTriExtensions) - 1
	}
	return fmt.Sprintf("%v%s", size, kMilBilTriExtensions[extensionIndex])
}

func Float64ByteSizeToString(size float64, limit int) string {
	extensionIndex := 0
	for noOfDigits(int(size))+len(byteSizeExtensions[extensionIndex])+float64Faction > limit && extensionIndex < len(byteSizeExtensions) {
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

func PostfixSpace(s string, limit int) string {
	for len(s) <= limit {
		s += " "
	}
	return s[:limit]
}

func PercentageToCharRep(rep string, count int, total int, limit int) string {
	number := int((float64(count) / float64(total)) * float64(limit))
	s := ""
	for len(s) < number {
		s += rep
	}
	return s
}

func noOfDigits(number int) int {
	if number == 0 {
		return 1
	}
	if number < 0 {
		number *= -1
	}
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}
