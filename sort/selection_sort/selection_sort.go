// Selection sort - O(n^2)
package selection_sort

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
