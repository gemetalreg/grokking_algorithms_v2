package dijkstra

import "math"

var graph = map[string]map[string]int{
	"start": {
		"A": 6,
		"B": 2,
	},
	"A": {
		"end": 1,
	},
	"B": {
		"A":   3,
		"end": 5,
	},
	"end": {},
}

var costs = map[string]int{
	"A":   6,
	"B":   2,
	"end": math.MaxInt,
}

// child - parents
var parents = map[string]string{
	"A":   "start",
	"B":   "start",
	"end": "",
}

var processes = make(map[string]struct{})
