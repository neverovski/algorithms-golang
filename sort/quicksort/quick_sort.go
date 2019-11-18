// Quick sort - O(n * log n)
package quicksort

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
