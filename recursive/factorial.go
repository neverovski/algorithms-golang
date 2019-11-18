// Factorial - n!
package recursive

func Factorial(el int) int {
	if el == 0 {
		return 1
	}

	return el * Factorial(el-1)
}
