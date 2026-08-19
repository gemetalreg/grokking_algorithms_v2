package sortchoose

import "testing"

func TestSortChoose(t *testing.T) {
	tests := []struct {
		unsorted []int
		sorted   []int
	}{
		{[]int{5, 3, 6, 2, 10}, []int{2, 3, 5, 6, 10}},
	}

	for _, test := range tests {
		sorted := selectionSort(test.unsorted)
		for i, el := range sorted {
			if el != test.sorted[i] {
				t.Errorf("wrong sort, i %d as el %d != %d from right sort", i, el, test.sorted[i])
				break
			}
		}

	}
}
