package fluent

import "testing"

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
