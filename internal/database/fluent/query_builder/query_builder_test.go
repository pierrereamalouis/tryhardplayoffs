package querybuilder

import "testing"

func TestQueryBuilder_Select(t *testing.T) {
	qb := new(QueryBuilder)

	qb.Select("name", "age").From("users")

	expected := "SELECT name, age FROM users "
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}
	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Select failed: expected %s, got %s", expected, actual)
	}
}

func TestQueryBuilder_SelectAll(t *testing.T) {
	qb := new(QueryBuilder)

	qb.SelectAllFrom("users")

	expected := "SELECT * FROM users "
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}
	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Select failed: expected %s, got %s", expected, actual)
	}
}

func TestQueryBuilder_Insert_Singular(t *testing.T) {
	qb := new(QueryBuilder)

	qb.InsertInto("employees").Columns("id", "name", "salary").Values(
		1, "John", 60000,
	)

	expected := "INSERT INTO employees (id, name, salary) VALUES ($1, $2, $3)"
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}
	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Insert failed: expected %s --- got %s", expected, actual)
	}
}

func TestQueryBuilder_Insert_Multiple(t *testing.T) {
	qb := new(QueryBuilder)

	qb.InsertInto("employees").Columns("id", "name", "salary").Values(
		1, "John", 60000,
		2, "Marie", 75000,
		3, "Luke", 80000,
	)

	expected := "INSERT INTO employees (id, name, salary) VALUES ($1, $2, $3), ($1, $2, $3), ($1, $2, $3)"
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}
	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Insert failed: expected %s --- got %s", expected, actual)
	}
}

func TestQueryBuilder_Update(t *testing.T) {
	qb := new(QueryBuilder)

	qb.Update("employees").Set(
		"id", 1,
		"name", "John",
		"salary", 60000,
	)

	expected := "UPDATE employees SET id = $1, name = $2, salary = $3"
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}
	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Insert failed: expected %s --- got %s", expected, actual)
	}
}

func TestQueryBuilder_Where(t *testing.T) {
	qb := new(QueryBuilder)
	qb.SelectAllFrom("users").Where("id", "=", "5")

	expected := "SELECT * FROM users WHERE id = $1"
	err := qb.Build()
	if err != nil {
		t.Errorf("TestQueryBuilder_Insert failed, could not build query : %s", err)
	}

	actual := qb.Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Where failed: expected %s --- got %s", expected, actual)
	}
}
