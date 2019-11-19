package breadth_first_search

import "testing"

var dataCases = []struct {
	in        map[string][]string
	startNode string
	endNode   string
	out       bool
	elements  int
}{
	{make(map[string][]string), "Frankfurt", "Munich", false, 0},
}

func TestBreadthFirstSearch(t *testing.T) {
	for _, value := range dataCases {
		in, elements := BreadthFirstSearch(value.in, value.startNode, value.endNode)

		if in != value.out && elements != value.elements {
			t.Errorf("ERROR: in - %v, out - %v, elemnts - %d", in, value.out, value.elements)
		}
	}
}
