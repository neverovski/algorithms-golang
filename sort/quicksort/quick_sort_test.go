package quicksort

import (
	"reflect"
	"testing"
)

var dataCases = []struct {
	in  []int
	out []int
}{
	{[]int{}, []int{}},
	{[]int{10, 20, 4, 5, 3, 1, 11}, []int{1, 3, 4, 5, 10, 11, 20}},
	{[]int{10}, []int{10}},
	{[]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	{[]int{2, 1}, []int{1, 2}},
}

func TestQSort(t *testing.T) {
	for n, value := range dataCases {
		in := QuickSort(value.in)
		out := value.out

		if !reflect.DeepEqual(in, out) {
			t.Errorf("Element number %d in %v out %v \n", n, in, out)
		}
	}
}
