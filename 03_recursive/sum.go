package main

import "fmt"

func main() {
	fmt.Println(Sum([]int{1, 3, 5, 6, 7})) // 22
}

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}
