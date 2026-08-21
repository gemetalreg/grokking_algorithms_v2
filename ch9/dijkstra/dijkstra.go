package dijkstra

import "math"

func dijkstra() {
	node := find_lowest_cost_node(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for neighbor, c := range neighbors {
			new_cost := cost + c
			if costs[neighbor] > new_cost {
				costs[neighbor] = new_cost
				parents[neighbor] = node
			}
		}
		processes[node] = struct{}{}
		node = find_lowest_cost_node(costs)
	}
}

func find_lowest_cost_node(costs map[string]int) string {
	lowest_cost := math.MaxInt
	lowest_cost_node := ""
	for node := range costs {
		cost := costs[node]
		if _, ok := processes[node]; !ok && cost < lowest_cost {
			lowest_cost = cost
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}
