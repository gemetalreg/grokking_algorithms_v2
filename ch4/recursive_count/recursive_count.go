package recursivecount

func count(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return 1 + count(slice[1:])
}
