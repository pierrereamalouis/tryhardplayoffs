package fluent

import (
	"fmt"
	"strings"
)

type (
	QueryType string
	Operator  string
	Chaining  string
	Join      string
	OrderBy   string

	QueryBuilder struct {
		Type         QueryType
		Fields       []string
		TableName    string
		WhereClauses []WhereClause
		JoinClauses  []JoinClause
		GroupBy      []string
		Having       []WhereClause
		OrderBy      []OrderByClause
		Limit        int
		Offset       int
		Final        string
	}

	WhereClause struct {
		Field      string
		Operator   Operator
		Value      any
		MultiValue []any // For cases like BETWEEN, IN
		Connector  Chaining
	}

	JoinClause struct {
		Type  Join
		Table string
		On    WhereClause
	}

	OrderByClause struct {
		Field     string
		Direction OrderBy
	}
)

const (
	// Query Type
	Select QueryType = "SELECT"
	Insert QueryType = "INSERT"
	Update QueryType = "UPDATE"
	Delete QueryType = "DELETE"

	// Comparison
	EQ   Operator = "="
	NEQ  Operator = "!="
	LT   Operator = "<"
	GT   Operator = ">"
	LTEQ Operator = "<="
	GTEQ Operator = ">="

	// Logical / Chaining
	And Chaining = "AND"
	Or  Chaining = "OR"
	Not Operator = "NOT"

	// Pattern Matching
	Like      Operator = "LIKE"
	ILike     Operator = "ILIKE"
	SimilarTo Operator = "SIMILAR TO"

	// Range
	Between    Operator = "BETWEEN"
	NotBetween Operator = "NOT BETWEEN"
	In         Operator = "IN"
	NotIn      Operator = "NOTIN"

	// Null Comparison
	IsNull    Operator = "IS NULL"
	IsNotNull Operator = "IS NOT NULL"

	// Array
	Any Operator = "ANY"
	All Operator = "ALL"

	// Join Clauses
	Inner Join = "INNER JOIN"
	Left  Join = "LEFT JOIN"
	Right Join = "RIGHT JOIN"
	Full  Join = "FULL"
	Cross Join = "CROSS JOIN"

	// Order By Clause
	ASC  OrderBy = "ASC"
	DESC OrderBy = "DESC"
)

func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	q.Type = Select
	q.Fields = fields

	return q
}

func (q *QueryBuilder) Insert(fields ...string) *QueryBuilder {
	q.Type = Insert
	q.Fields = fields

	return q
}

func (q *QueryBuilder) Update(fields ...string) *QueryBuilder {
	q.Type = Update
	q.Fields = fields

	return q
}

func (q *QueryBuilder) Delete() *QueryBuilder {
	q.Type = Delete
	return q
}

func (q *QueryBuilder) Table(table string) *QueryBuilder {
	q.TableName = table

	return q
}

func (q *QueryBuilder) Where(field string, operator Operator, value any, values ...any) *QueryBuilder {
	clause := WhereClause{Field: field, Operator: operator, Value: value}

	if len(q.WhereClauses) > 0 {
		clause.Connector = And
	}

	if values != nil {
		clause.MultiValue = values
	}

	q.WhereClauses = append(q.WhereClauses, clause)

	return q
}

func (q *QueryBuilder) AndWhere(field string, operator Operator, value any, values ...any) *QueryBuilder {
	clause := WhereClause{Field: field, Operator: operator, Value: value, Connector: And}

	if values != nil {
		clause.MultiValue = values
	}

	q.WhereClauses = append(q.WhereClauses, clause)
	return q
}

func (q *QueryBuilder) OrWhere(field string, operator Operator, value any, values ...any) *QueryBuilder {
	clause := WhereClause{Field: field, Operator: operator, Value: value, Connector: And}

	if values != nil {
		clause.MultiValue = values
	}

	q.WhereClauses = append(q.WhereClauses, clause)

	return q
}

func (q *QueryBuilder) WhereBetween(field string, value1 any, value2 any) *QueryBuilder {
	values := []any{value1, value2}
	q.WhereClauses = append(q.WhereClauses, WhereClause{Field: field, Operator: Between, MultiValue: values})

	return q
}

func (q *QueryBuilder) WhereIn(field string, values ...any) *QueryBuilder {
	q.WhereClauses = append(q.WhereClauses, WhereClause{Field: field, Operator: In, MultiValue: values})
	return q
}

func (q *QueryBuilder) WhereNotBetween(field string, values ...any) *QueryBuilder {
	q.WhereClauses = append(q.WhereClauses, WhereClause{Field: field, Operator: NotBetween, MultiValue: values})

	return q
}

func (q *QueryBuilder) WhereNotIn(field string, values ...any) *QueryBuilder {
	q.WhereClauses = append(q.WhereClauses, WhereClause{Field: field, Operator: NotIn, MultiValue: values})
	return q
}

func (q *QueryBuilder) Build() *QueryBuilder {

	q.Final += q.buildOpenStatement()

	if len(q.WhereClauses) > 0 {
		q.buildWhereClause()
	}

	return q
}

func (q *QueryBuilder) buildOpenStatement() string {
	var openStatement string

	switch q.Type {
	case "SELECT":
		openStatement = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(q.Fields, ", "), q.TableName)
	case "INSERT":
		openStatement = fmt.Sprintf("INSERT INTO %s (%s) VALUES ", q.TableName, strings.Join(q.Fields, ", "))

	case "UPDATE":
		openStatement = fmt.Sprintf("UPDATE %s SET %s", q.TableName, strings.Join(q.Fields, " = ?, "))
	case "DELETE":
		openStatement = fmt.Sprintf("DELETE FROM %s ", q.TableName)
	}

	return openStatement
}

func (q *QueryBuilder) buildWhereClause() string {
	var whereClause string

	return whereClause
}

func (q *QueryBuilder) builderInsertValues() string {
	var insertValues string

	return insertValues
}
