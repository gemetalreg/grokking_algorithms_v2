package sortchoose

import "slices"

func findSmallest(slice []int) int {
	smallest_el := slice[0]
	smallest_index := 0
	for i, el := range slice[1:] {
		real_i := i + 1
		if el < smallest_el {
			smallest_index = real_i
		}
	}
	return smallest_index
}

func selectionSort(src []int) []int {
	newSlice := []int{}
	dst := make([]int, len(src))
	copy(dst, src)
	for range src {
		smallest_i := findSmallest(dst)
		newSlice = append(newSlice, dst[smallest_i])
		dst = slices.Delete(dst, smallest_i, smallest_i+1)
	}
	return newSlice
}
