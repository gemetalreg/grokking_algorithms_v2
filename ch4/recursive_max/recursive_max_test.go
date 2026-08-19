package recursivemax

import "testing"

func TestRecursiveMax(t *testing.T) {
	tests := []struct {
		arr []int
		m   int
	}{
		{[]int{1, 23, 4, 6}, 23},
	}

	for _, test := range tests {
		if max(test.arr) != test.m {
			t.Errorf("%d must be max el", test.m)
		}
	}
}
