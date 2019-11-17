// Factorial - n!
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorial(5)) // 120
}

func Factorial(el int) int {
	if el == 0 {
		return 1
	}

	return el * Factorial(el-1)
}
