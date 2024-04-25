package fluent

/*** Open Statement Delimiter ***/
// SELECT
func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	q.Type = Select

	q.setFields(fields...)

	return q
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.TableName = table

	return q
}

func (q *QueryBuilder) SelectAllFrom(table string) *QueryBuilder {
	q.SelectAll = true

	return q
}

// INSERT
func (q *QueryBuilder) InsertInto(table string) *QueryBuilder {
	q.Type = Insert
	q.table(table)

	return q
}

func (q *QueryBuilder) Columns(fields ...string) *QueryBuilder {
	if q.Type != Insert {
		// TODO : refactor to better handler error
		panic("Values method can only be used with INSERT queries")
	}

	q.Type = Insert

	q.setFields(fields...)

	return q
}

func (q *QueryBuilder) Values(values ...any) *QueryBuilder {
	if q.Type != Insert {
		// TODO : refactor to better handler error
		panic("Values method can only be used with INSERT queries")
	}

	// Check if the number of values is a multiple of the number of fields
	if len(values)%len(q.Fields) != 0 {
		panic("Number of values must be a multiple of the number of fields")
	}

	// Calculate the number of rows based on the number of fields
	numRows := len(values) / len(q.Fields)

	q.InsertValues = make([]map[string]any, numRows)

	/* create a slice of map[string]any
	 if for example q.Fields have ["id", "name", "salary"]
	 need to have a slice of the values pass Values (
		1, "John", 60000,
		2, "Marie", 75000,
		3, "Luke", 80000
	 )

	 q.InsertValue = [{id: 1, name: "John", 60000}, {id: 2, "Marie", 75000}, {id: 3, "Luke", 80000} ]
	*/

	for i := 0; i < numRows; i++ {
		// Create a map to hold the values for the current row
		row := make(map[string]any)

		// Iterate over the fields and assign the corresponding value
		for j, field := range q.Fields {
			// Calculate the index of the value for the current field
			idx := i*len(q.Fields) + j

			// Assign the vlaue to the field in the map

			row[field] = values[idx]
		}
	}

	return q
}

// UPDATE
func (q *QueryBuilder) Update(table string) *QueryBuilder {
	q.Type = Update
	q.TableName = table

	return q
}

func (q *QueryBuilder) Set(values ...any) *QueryBuilder {
	if q.Type != Update {
		// TODO : refactor to better handler error
		panic("Values method can only be used with INSERT queries")
	}

	if len(values)%2 != 0 {
		panic("Must have key:value pair matches")
	}

	// Create a map to hold the values of current set key:value pair
	setKeyValue := make(map[string]any)

	for i := 0; i < len(values); i += 2 {
		field, ok := values[i].(string)

		if !ok {
			panic("Key must be string")
		}

		value := values[i+1]

		setKeyValue[field] = value
	}

	q.UpdateValues = setKeyValue

	return q
}

// DELETE
func (q *QueryBuilder) DeleteFrom() *QueryBuilder {
	q.Type = Delete
	return q
}

/*** End Open Statement Delimiter ***/
