package main

import (
	"testing"
)

var dataBinarySearch = []struct {
	in  []int
	el  int
	out bool
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10, false},
	{[]int{}, 10, false},
	{[]int{20}, 20, true},
	{[]int{10, 20}, 20, true},
	{[]int{10}, 5, false},
	{[]int{1, 2, 3, 4, 5, 6, 7, 10}, 10, true},
}

func TestBinarySearch(t *testing.T) {
	for n, value := range dataBinarySearch {
		in := BinarySearch(value.in, value.el)
		out := value.out

		if in != out {
			t.Errorf("Element number %d in %t out %t \n", n, in, out)
		}
	}
}
