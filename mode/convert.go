package mode

import "math"

func RoundToTwoDecimal(value float64) float64 {
	return math.Round(value*100) / 100
}

func ConvertMemory(bytes float64) (float64, string) {
	if bytes >= 1024*1024*1024 {
		return RoundToTwoDecimal(bytes / (1024 * 1024 * 1024)), "GB"
	} else {
		return RoundToTwoDecimal(bytes / (1024 * 1024)), "MB"
	}
}

func ConvertDisk(bytes float64) float64 {
	return RoundToTwoDecimal(bytes / (1024 * 1024 * 1024))
}

func ConvertNetwork(bytes float64) (float64, string) {
	if bytes >= 1000*1024 {
		return RoundToTwoDecimal(bytes / (1024 * 1024)), "MB"
	} else {
		return RoundToTwoDecimal(bytes / 1024), "kB"
	}
}
