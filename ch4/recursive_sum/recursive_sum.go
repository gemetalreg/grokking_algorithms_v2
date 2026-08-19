package recursivesum

func sum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return slice[0] + sum(slice[1:])
}
