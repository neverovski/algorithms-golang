package breadth_first_search

func BreadthFirstSearch(graph map[string][]string, startNode string, endNode string) (bool, int) {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[startNode]...)
	var searched []string
	var pr []string
	pr = append(pr, startNode)

	for len(searchQueue) != 0 {
		var node = searchQueue[0]
		searchQueue = searchQueue[1:]
		nodeAlreadySearched := false

		for i := 0; i < len(searched); i++ {
			nodeAlreadySearched = searched[i] == node
		}

		if !nodeAlreadySearched {
			pr = append(pr, node)
			if node == endNode {
				break
			}
			searchQueue = append(searchQueue, graph[node]...)
			searched = append(searched, node)
		}
	}

	return false, 0
}
