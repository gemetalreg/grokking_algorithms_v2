package graph

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func BFS(dirname string) {
	queue := Queue[string]{}
	queue.Enqueue(dirname)
	for !queue.IsEmpty() {
		dir, _ := queue.Dequeue()
		entries, err := os.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, entry := range entries {
			full := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				queue.Enqueue(full)
			} else {
				fmt.Println(full)
			}
		}
	}
}