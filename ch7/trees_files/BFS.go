package trees_files

import (
	"log"
	"os"
	"path/filepath"
)

func BFS(dirname string) []string {
	queue := Queue[string]{}
	queue.Enqueue(dirname)
	files := Queue[string]{}
	for !queue.IsEmpty() {
		dir, _ := queue.Dequeue()
		entries, err := os.ReadDir(dir)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, entry := range entries {
			full := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				queue.Enqueue(full)
			} else {
				files.Enqueue(full)
			}
		}
	}
	return files.Elements()
}
