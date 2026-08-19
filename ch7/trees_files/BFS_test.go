package trees_files

import (
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {

	files := BFS(".")
	slices.Sort(files)
	ans := []string{"BFS_test.go", "BFS.go", "go.mod", "queue.go"}
	slices.Sort(ans)
	for i, file := range files {
		if file != ans[i] {
			t.Errorf("%s is not equal %s in sorted slices", file, ans[i])
		}
	}
}
