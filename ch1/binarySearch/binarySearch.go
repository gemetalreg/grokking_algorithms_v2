package bs

import "fmt"

func BinarySearch(slice []int, item int) (int, error) {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]

		if guess == item {
			return mid, nil
		} else if guess < item {
			low = mid + 1
		} else if guess > item {
			high = mid - 1
		}
	}
	return 0, fmt.Errorf("no %d in slice", item)
}
