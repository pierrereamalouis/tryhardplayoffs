package fluent

import (
	"fmt"
	"testing"
	"time"
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

func TestCollection_InsertQueryString(t *testing.T) {
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
	queryString := collection.QB.Final

	fmt.Println(queryString)

	if queryString != expected {
		t.Errorf("TestCollection_InsertQueryString failed: expected %s, got %s", expected, queryString)
	}
}

func TestCollection_Insert(t *testing.T) {
	fluent := new(Fluent)

	// TODO : init container to get pgxpool instance
	collection := fluent.Collection("nhl_teams")

	nhlTeam := map[string]any{
		"city":         "MontreaL",
		"name":         "Canadiens",
		"abbreviation": "MTL",
		"created_on":   time.Now(),
	}

	_, err := collection.Insert(nhlTeam)

	if err != nil {
		t.Errorf("TestCollection_Insert failed: %s", err)
	}
}
