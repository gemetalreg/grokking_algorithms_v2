package bs

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		slice []int
		item  int
		want  bool
	}{
		{[]int{1, 3, 5, 7, 9}, 3, true},
		{[]int{1, 3, 5, 7, 9}, -1, false},
	}

	for _, test := range tests {
		if _, err := BinarySearch(test.slice, test.item); (err != nil) == test.want {
			t.Errorf("%d not in slice\n", test.item)
		}
	}
}
