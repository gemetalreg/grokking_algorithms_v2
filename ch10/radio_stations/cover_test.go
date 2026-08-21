package radiostations

import (
	"testing"
)

func TestCover(t *testing.T) {
	Cover()
	ans := set{
		"ktwo":   struct{}{},
		"kthree": struct{}{},
		"kone":   struct{}{},
		"kfive":  struct{}{},
	}

	if len(*(ans.Intersect(&final_stations))) != 4 {
		t.Error("wrong final stations")
	}
}
