package graph

func BFS(name string) string {
	graph := CreateGraph()
	queue := Queue[string]{}
	queue.Enqueue(name)
	set := make(map[string]struct{})
	for !queue.IsEmpty() {
		person, _ := queue.Dequeue()
		if _, ok := set[person]; !ok {
			if person_is_seller(person) {
				return person
			}
		}

		for _, p := range graph[person] {
			queue.Enqueue(p)
		}
		set[person] = struct{}{}
	}
	return ""
}

func person_is_seller(name string) bool {
	runes := []rune(name)
	return runes[len(runes)-1] == 'm'
}
