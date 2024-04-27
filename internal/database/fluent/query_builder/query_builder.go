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
		Type             QueryType
		Fields           []string
		SelectAll        bool
		TableName        string
		InsertValues     []map[string]any
		UpdateSetValues  map[string]any
		UpdateSetFields  []string
		WhereClauses     []WhereClause
		JoinClauses      []JoinClause
		GroupBy          []string
		Having           []WhereClause
		OrderBy          []OrderByClause
		Limit            int
		Offset           int
		Final            string
		PlaceholderIndex int
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

func (q *QueryBuilder) setFields(fields ...string) *QueryBuilder {
	q.Fields = fields

	return q
}

func (q *QueryBuilder) table(table string) *QueryBuilder {
	q.TableName = table

	return q
}

// TODO : implement reset func to remove all data once query is executed and completed
func (q *QueryBuilder) reset() {

}

func (q *QueryBuilder) Build() *QueryBuilder {

	q.Final += q.buildOpenStatement()

	if len(q.WhereClauses) > 0 {
		q.buildWhereClause()
	}

	q.reset()

	// don't need to return q as it would be empty after q.reset()
	return q
}

func (q *QueryBuilder) buildOpenStatement() string {
	var openStatement string

	switch q.Type {
	case "SELECT":
		if q.SelectAll {
			openStatement = fmt.Sprintf("SELECT * FROM %s ", q.TableName)
		} else {
			openStatement = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(q.Fields, ", "), q.TableName)
		}
	case "INSERT":
		openStatement = fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", q.TableName, strings.Join(q.Fields, ", "), q.buildInsertValues())
	case "UPDATE":
		openStatement = fmt.Sprintf("UPDATE %s SET %s", q.TableName, q.buildUpdateValues())
	case "DELETE":
		openStatement = fmt.Sprintf("DELETE FROM %s ", q.TableName)
	}

	return openStatement
}

// help handle singular insert or multiple insert in one query
// iterate through the InsertValues slice in order
// to build the string after "VALUES" of the "INSERT" statement
// final string:
// multiple insert string "($1, $2, $3), ($1, $2, $3), ($1, $2, $3)"
// singular  string "($1, $2, $3)"
func (q *QueryBuilder) buildInsertValues() string {
	var valuesStatement []string

	for range q.InsertValues {

		// create temporary slice of string
		var rowFields []string

		for index := range q.Fields {
			q.PlaceholderIndex = index + 1
			rowFields = append(rowFields, fmt.Sprintf("$%d", q.PlaceholderIndex))
		}

		valuesStatement = append(valuesStatement, fmt.Sprintf("(%s)", strings.Join(rowFields, ", ")))
	}

	// final string "($1, $2, $3), ($1, $2, $3), ($1, $2, $3)"
	return strings.Join(valuesStatement, ", ")
}

func (q *QueryBuilder) buildUpdateValues() string {
	// q.InsertValues = []interface{
	// 	{1, "John", 60000},
	// }

	// [string]{" name = $1, age = $2 "}

	var setValuesClause []string
	// var index = 1

	// for key := range q.UpdateSetValues {

	// 	setValuesClause = append(setValuesClause, fmt.Sprintf("%s = $%d", key, index))

	// 	index++
	// }

	for index, field := range q.UpdateSetFields {
		q.PlaceholderIndex = index + 1

		setValuesClause = append(setValuesClause, fmt.Sprintf("%s = $%d", field, q.PlaceholderIndex))
	}

	return strings.Join(setValuesClause, ", ")
}

func (q *QueryBuilder) buildWhereClause() string {
	var whereClauses []string
	// var rangeOperators = []string{"BETWEEN", "NOT BETWEEN", "NOT IN", "IN"}

	/* WhereClause struct {
		Field      string
		Operator   Operator
		Value      any
		MultiValue []any // For cases like BETWEEN, IN
		Connector  Chaining
	}*/

	for _, clause := range q.WhereClauses {
		switch clause.Operator {
		case Between:
			fallthrough
		case NotBetween:
			placeHolderOne := q.PlaceholderIndex + 1
			placeHolderTwo := placeHolderOne + 1
			var clauseStr string

			if clause.Connector != "" {
				clauseStr = fmt.Sprintf("%s %s %s $%d AND $%d", clause.Connector, clause.Field, clause.Operator, placeHolderOne, placeHolderTwo)
			} else {
				clauseStr = fmt.Sprintf("%s %s $%d AND $%d", clause.Field, clause.Operator, placeHolderOne, placeHolderTwo)
			}

			whereClauses = append(whereClauses, clauseStr)
		case In:
			whereClauses = append(whereClauses, fmt.Sprintf("%s IN (%s)", clause.Field, q.buildWhereInClause()))
		}
		// if rangeOperators. clause.Operator

		// whereClause += fmt.Sprintf("%s %s ")
	}

	return strings.Join(whereClauses, " ")
}

func (q *QueryBuilder) buildWhereInClause() string {
	return ""
}
