package main

import (
	SPF "algorithms-golang/dijkstra_algorithm"
	hash "algorithms-golang/hash_tables"
	"algorithms-golang/recursive"
	BFS "algorithms-golang/search/breadth_first_search"
	"algorithms-golang/sort/quicksort"
	"algorithms-golang/sort/selection_sort"
	"fmt"
)

func main() {
	hashTables()

	fmt.Println(recursive.Factorial(5))              // 120
	fmt.Println(recursive.Fibonacci(10))             // 55
	fmt.Println(recursive.Sum([]int{1, 3, 5, 6, 7})) // 22

	graph := make(map[string][]string)
	graph["Frankfurt"] = []string{"Mannheim", "Würzburg", "Kassel"}
	graph["Mannheim"] = []string{"Karlsruhe"}
	graph["Karlsruhe"] = []string{"Augsburg"}
	graph["Augsburg"] = []string{}
	graph["Würzburg"] = []string{"Nuremberg", "Erfurt"}
	graph["Nuremberg"] = []string{"Stuttgart"}
	graph["Erfurt"] = []string{}
	graph["Kassel"] = []string{"Munich"}
	graph["Munich"] = []string{}
	fmt.Println(BFS.BreadthFirstSearch(graph, "Frankfurt", "Munich"))

	graphDistance := make(map[string]map[string]int)
	graphDistance["Frankfurt"] = map[string]int{}
	graphDistance["Frankfurt"]["Mannheim"] = 100
	graphDistance["Frankfurt"]["Würzburg"] = 150
	graphDistance["Frankfurt"]["Kassel"] = 90

	graphDistance["Mannheim"] = map[string]int{}
	graphDistance["Mannheim"]["Munich"] = 100

	graphDistance["Würzburg"] = map[string]int{}
	graphDistance["Würzburg"]["Munich"] = 130

	graphDistance["Kassel"] = map[string]int{}
	graphDistance["Kassel"]["Munich"] = 140

	graphDistance["Munich"] = map[string]int{}
	fmt.Println(SPF.DijkstraAlgorithm(graphDistance, "Frankfurt", "Munich")) // ["Frankfurt","Mannheim","Munich"]

	fmt.Println(quicksort.QuickSort([]int{10, 4, 5, 12, 1, 6}))          // [1,4,5,6,10]
	fmt.Println(selection_sort.SelectionSort([]int{10, 2, 4, 1, 8, 11})) // [1,2,4,8,10,11]
}

// Hash tables
func hashTables() {
	fmt.Println("SetData - Tom:", hash.SetData("Tom", 30))     // true
	fmt.Println("SetData - Harry:", hash.SetData("Harry", 25)) // true

	fmt.Println(hash.GetData("Tom")) // (30,true)
	fmt.Println(hash.GetData("TOM")) // (0,false)

	fmt.Println("DeleteData - Harry:", hash.DeleteData("Harry")) // true
	fmt.Println("DeleteData - Harry:", hash.DeleteData("Harry")) // false
}
