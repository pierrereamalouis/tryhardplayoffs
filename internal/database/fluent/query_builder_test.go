package fluent

import "testing"

func TestQueryBuilder_Select(t *testing.T) {
	qb := new(QueryBuilder)

	qb.Table("users").Select("name", "age")

	expected := "SELECT name, age FROM users "
	actual := qb.Build().Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Select failed: expected %s, got %s", expected, actual)
	}
}

func TestQueryBuilder_Insert(t *testing.T) {
	qb := new(QueryBuilder)

	qb.Table("employees").Insert("id", "name", "salary")

	expected := "INSERT INTO employees (id, name, salary)"
	actual := qb.Build().Final

	if actual != expected {
		t.Errorf("TestQueryBuilder_Insert failed: expected %s, got %s", expected, actual)
	}
}
