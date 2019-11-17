// Selection sort - O(n^2)
package main

import "fmt"

func main() {

	fmt.Println(SelectionSort([]int{10, 2, 4, 1, 8, 11})) // [1,2,4,8,10,11]
}

func SelectionSort(list []int) []int {
	size := len(list)
	newList := make([]int, size)

	for i := 0; i < size; i++ {
		index := findSmallestIndex(list)
		newList[i] = list[index]
		list = append(list[:index], list[index+1:]...)
	}

	return newList
}

func findSmallestIndex(list []int) int {
	smallElement := list[0]
	smallIndex := 0

	for i := 1; i < len(list); i++ {
		if list[i] < smallElement {
			smallElement = list[i]
			smallIndex = i
		}
	}

	return smallIndex
}
