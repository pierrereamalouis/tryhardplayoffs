package fluent

import (
	"fmt"
	"testing"
)

type Employee struct {
	id     int
	name   string
	salary int
}

func TestCollection_Find(t *testing.T) {
	fluent := new(Fluent)

	collection := fluent.Collection("users")

	collection.Find(Cond{"id =": 5})

	expected := "SELECT * FROM users WHERE id = $1"
	queryString := collection.Build()

	if queryString != expected {
		t.Errorf("TestCollection_Find failed: expected %s, got %s", expected, queryString)
	}
}

func TestCollection_Insert(t *testing.T) {
	fluent := new(Fluent)

	collection := fluent.Collection("employees")

	// employeeStruct := Employee{id: 1, name: "Sally", salary: 75000}
	employee := map[string]any{
		"id":     5,
		"name":   "Sally",
		"salary": 75000,
	}

	collection.Insert(employee)

	expected := "INSERT INTO employees (id, name, salary) VALUES ($1, $2, $3)"
	queryString := collection.Build()

	fmt.Println(queryString)

	if queryString != expected {
		t.Errorf("TestCollection_Insert failed: expected %s, got %s", expected, queryString)
	}
}
