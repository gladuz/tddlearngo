package arraysslices

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(numbers []int) int {
	sum := 0
	for _, x := range numbers {
		sum += x
	}
	return sum
}

func SumZip(sliceZip ...[]int) []int {
	min_len := math.MaxInt
	for _, sl := range sliceZip {
		min_len = min(len(sl), min_len)
	}
	sum_slice := make([]int, min_len)
	for i := 0; i < min_len; i++ {
		sum := 0
		for _, x := range sliceZip {
			sum += x[i]
		}
		sum_slice[i] = sum
	}
	return sum_slice
}
