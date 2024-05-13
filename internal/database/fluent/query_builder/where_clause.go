package querybuilder

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
