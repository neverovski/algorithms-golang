// Quick sort - O(n * log n)
package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{10, 4, 5, 12, 1, 6})) // [1,4,5,6,10]
}

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]

	var less []int
	var greater []int
	for _, num := range list[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = append(QuickSort(less), pivot)
	greater = QuickSort(greater)
	return append(less, greater...)
}
