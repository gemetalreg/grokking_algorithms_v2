package graph

import "testing"

func TestBFS(t *testing.T) {
	tests := []struct {
		name        string
		seller_name string
	}{
		{"you", "Tom"},
	}

	for _, test := range tests {
		if BFS(test.name) != test.seller_name {
			t.Errorf("name %q has no way to mango seller", test.name)
		}
	}
}
