package selection_sort

import (
	"reflect"
	"testing"
)

var dataCases = []struct {
	in  []int
	out []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{3, 1}, []int{1, 3}},
	{[]int{10, 4, 5, 19, 6, 7, 11, 3, 1}, []int{1, 3, 4, 5, 6, 7, 10, 11, 19}},
}

func TestSelectionSort(t *testing.T) {
	for n, value := range dataCases {
		in := SelectionSort(value.in)
		out := value.out

		if !reflect.DeepEqual(in, out) {
			t.Errorf("Element number %d in %v out %v \n", n, in, out)
		}
	}
}
