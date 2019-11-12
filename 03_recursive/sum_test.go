package main

import "testing"

var dataSum = []struct {
	in  []int
	out int
}{
	{[]int{}, 0},
	{[]int{1}, 1},
	{[]int{0, 8, 10, 20, 2, 3, 4, 5}, 52},
}

func TestSum(t *testing.T) {
	for n, value := range dataSum {
		in := Sum(value.in)
		out := value.out

		if in != out {
			t.Errorf("Element number %d in %d out %d \n", n, in, out)
		}
	}
}
