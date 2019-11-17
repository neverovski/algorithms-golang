// Binary search - O(log n)
package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 5, 6, 7, 8, 10, 11}

	fmt.Println(BinarySearch(list, 10)) // true
	fmt.Println(BinarySearch(list, 20)) // false
}

func BinarySearch(list []int, el int) bool {
	var iteration int
	var low int
	high := len(list) - 1

	for low <= high {
		iteration++
		mid := (low + high) / 2

		if el == list[mid] {
			return true
		}

		if el < list[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
