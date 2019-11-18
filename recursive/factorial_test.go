package recursive

import "testing"

var dataFactorial = []struct {
	in  int
	out int
}{
	{0, 1},
	{1, 1},
	{5, 120},
}

func TestFactorial(t *testing.T) {
	for n, value := range dataFactorial {
		in := Factorial(value.in)
		out := value.out

		if in != out {
			t.Errorf("Element number %d in %d out %d \n", n, in, out)
		}
	}
}
