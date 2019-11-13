package main

import "testing"

var dataHashTables = []struct {
	name string
	age  int
}{
	{"Tom", 30},
	{"", 15},
	{"Harry", 10},
	{"Races", 35},
	{"Adrian", 20},
}

func TestGetData(t *testing.T) {
	for _, data := range dataHashTables {
		setData := SetData(data.name, data.age)
		if setData {
			out := data.age

			if in, _ := GetData(data.name); out != in {
				t.Errorf("Error in %d out %d \n", in, out)
			}
		}
	}

	if _, out := GetData("TestCase"); out {
		t.Error("Error - key false")
	}
}

func TestDeleteData(t *testing.T) {
	for _, data := range dataHashTables {
		setData := SetData(data.name, data.age)
		if setData {

			if !DeleteData(data.name) {
				t.Error("Error key")
			}
		}
	}

	if DeleteData("TestCase") {
		t.Error("Error - key false")
	}
}
