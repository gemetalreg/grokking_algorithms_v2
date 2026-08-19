package graph

type graph_type map[string][]string

func CreateGraph() graph_type {
	graph := make(graph_type)
	graph["you"] = []string{"Alice", "Bob", "Kler"}
	graph["Bob"] = []string{"Anudz", "Peggi"}
	graph["Kler"] = []string{"Tom", "Jonny"}
	graph["Anudz"] = []string{}
	graph["Peggi"] = []string{}
	graph["Tom"] = []string{}
	graph["Jonny"] = []string{}

	return graph
}
