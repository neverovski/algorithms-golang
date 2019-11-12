// Fibonacci -  f(n)=f(n-1) + f(n-2)
package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(10)) // 55
}

func Fibonacci(el int) int {
	if el == 0 || el == 1 {
		return el
	}

	return Fibonacci(el-1) + Fibonacci(el-2)
}
