package recursivemax

func max(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	if len(slice) == 1 {
		return slice[0]
	}
	if len(slice) == 2 {
		if slice[0] > slice[1] {
			return slice[0]
		}
		return slice[1]
	}
	if slice[0] > max(slice[1:]) {
		return slice[0]
	}
	return max(slice[1:])
}
