package t

import (
	"math"
	"strconv"
)

const (
	// KB represents the size of a kilobyte.
	KB float64 = 1024
	// MB represents the size of a megabyte.
	MB = 1024 * KB
	// GB represents the size of a gigabyte.
	GB = 1024 * MB
	// TB represents the size of a terabyte.
	TB = 1024 * GB
	// PB represents the size of a petabyte.
	PB = 1024 * TB
)

func FileSize(bytesize int64, sigfig ...int) string {
	size := float64(bytesize)
	var sigfigs int
	if len(sigfig) == 0 {
		sigfigs = 2
	} else {
		sigfigs = sigfig[0]
	}
	switch {
	case size >= PB:
		return strconv.FormatFloat(math.Round((size*100)/PB)/100, 'f', sigfigs, 64) + "PB"
	case size >= TB:
		return strconv.FormatFloat(math.Round((size*100)/TB)/100, 'f', sigfigs, 64) + "TB"
	case size >= GB:
		return strconv.FormatFloat(math.Round((size*100)/GB)/100, 'f', sigfigs, 64) + "GB"
	case size >= MB:
		return strconv.FormatFloat(math.Round((size*100)/MB)/100, 'f', sigfigs, 64) + "MB"
	case size >= KB:
		return strconv.FormatFloat(math.Round((size*100)/KB)/100, 'f', sigfigs, 64) + "KB"
	default:
		return strconv.FormatInt(bytesize, 10) + "B"
	}
}
