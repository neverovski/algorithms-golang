// Hash tables - O(1)
package hash_tables

var data = map[string]int{}

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
