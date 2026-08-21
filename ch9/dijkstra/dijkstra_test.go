package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	dijkstra()
	if costs["end"] != 6 {
		t.Errorf("cost end must be 6, not %d", costs["end"])
	}
}
