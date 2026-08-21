package trees_files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func DFS(dirname string) {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Println(err)
		return
	}
	for _, entry := range entries {
		full := filepath.Join(dirname, entry.Name())
		if entry.IsDir() {
			DFS(full)
		} else {
			fmt.Println(full)
		}
	}
}
