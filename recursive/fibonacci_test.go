package main

import "testing"

var dataFibonacci = []struct {
	in  int
	out int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{10, 55},
	{12, 144},
}

func TestFibonacci(t *testing.T) {
	for n, value := range dataFibonacci {
		in := Fibonacci(value.in)
		out := value.out

		if in != out {
			t.Errorf("Element number %d in %d out %d \n", n, in, out)
		}
	}
}
