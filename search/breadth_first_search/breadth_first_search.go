package breadth_first_search

func breadthFirstSearch(graph map[string][]string, startNode string) {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[startNode]...)
	//var searched []string
}
