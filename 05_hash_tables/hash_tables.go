// Hash tables - O(1)
package main

import "fmt"

var data = map[string]int{}

func main() {
	fmt.Println("SetData - Tom:", SetData("Tom", 30))     // true
	fmt.Println("SetData - Harry:", SetData("Harry", 25)) // true

	fmt.Println(GetData("Tom")) // (30,true)
	fmt.Println(GetData("TOM")) // (0,false)

	fmt.Println("DeleteData - Harry:", DeleteData("Harry")) // true
	fmt.Println("DeleteData - Harry:", DeleteData("Harry")) // false
}

// Get data
func GetData(name string) (int, bool) {
	if age, ok := data[name]; ok {
		return age, ok
	}

	return 0, false
}

// Set data
func SetData(name string, age int) bool {
	if name != "" {
		data[name] = age
		return true
	}

	return false
}

// Delete data
func DeleteData(name string) bool {
	if _, ok := data[name]; ok {
		delete(data, name)
		return ok
	}

	return false
}
