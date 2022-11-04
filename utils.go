package typeutils

import (
	"fmt"
)

type CanContain interface {
	comparable
}

type BasicNums interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Contains[T CanContain](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ToByteFormat[T BasicNums](num T) string {
	number := float64(num)
	var unit string
	switch {
	case number >= 1024*1024*1024*1024:
		number /= 1024 * 1024 * 1024 * 1024
		unit = "TB"
	case number >= 1024*1024*1024:
		number /= 1024 * 1024 * 1024
		unit = "GB"
	case number >= 1024*1024:
		number /= 1024 * 1024
		unit = "MB"
	case number >= 1024:
		number /= 1024
		unit = "KB"
	default:
		unit = "B"
	}
	return fmt.Sprintf("%.2f %s", number, unit)
}

func ChunkSlice[T CanContain](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
