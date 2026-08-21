package radiostations

type set map[string]struct{}

func (s1 *set) Intersect(s2 *set) *set {
	result := make(set)

	smaller, bigger := s1, s2
	if len(*s1) > len(*s2) {
		smaller, bigger = s2, s1
	}

	for k := range *smaller {
		if _, ok := (*bigger)[k]; ok {
			result[k] = struct{}{}
		}
	}
	return &result
}

var states_needed = set{
	"mt": struct{}{},
	"wa": struct{}{},
	"or": struct{}{},
	"id": struct{}{},
	"nv": struct{}{},
	"ut": struct{}{},
	"ca": struct{}{},
	"az": struct{}{},
}

var stations = map[string]set{
	"kone": {
		"id": struct{}{},
		"nv": struct{}{},
		"ut": struct{}{},
	},
	"ktwo": {
		"wa": struct{}{},
		"id": struct{}{},
		"mt": struct{}{},
	},
	"kthree": {
		"or": struct{}{},
		"nv": struct{}{},
		"ca": struct{}{},
	},
	"kfour": {
		"nv": struct{}{},
		"ut": struct{}{},
	},
	"kfive": {
		"ca": struct{}{},
		"az": struct{}{},
	},
}

var final_stations = make(set)

