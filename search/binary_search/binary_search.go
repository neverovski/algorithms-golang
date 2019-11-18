// Binary search - O(log n)
package binary_search

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
