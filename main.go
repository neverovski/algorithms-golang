package main

import (
	hash "algorithms-golang/hash_tables"
	"algorithms-golang/recursive"
	"algorithms-golang/search/binary_search"
	"algorithms-golang/sort/quicksort"
	"algorithms-golang/sort/selection_sort"
	"fmt"
)

func main() {
	hashTables()

	fmt.Println(recursive.Factorial(5))              // 120
	fmt.Println(recursive.Fibonacci(10))             // 55
	fmt.Println(recursive.Sum([]int{1, 3, 5, 6, 7})) // 22

	list := []int{1, 2, 3, 5, 6, 7, 8, 10, 11}
	fmt.Println(binary_search.BinarySearch(list, 10)) // true
	fmt.Println(binary_search.BinarySearch(list, 20)) // false

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
