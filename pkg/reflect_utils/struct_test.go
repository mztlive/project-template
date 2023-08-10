package reflect_utils

import "testing"

func TestGetNameFromStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	name := GetNameFromStruct(p)

	if name != "Person" {
		t.Errorf("Expected name to be 'Person', but got '%s'", name)
	}
}

func TestGetSnakeNameFromStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	name := GetSnakeNameFromStruct(p)

	if name != "person" {
		t.Errorf("Expected name to be 'person', but got '%s'", name)
	}
}
