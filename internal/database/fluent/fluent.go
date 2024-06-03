package fluent

import (
	"context"
	"errors"
	"fmt"
	"strings"
	qb "tryhardplayoffs/internal/database/fluent/query_builder"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Fluent struct {
		DB  *pgxpool.Pool
		Col Collection
	}

	Collection struct {
		Fluent    *Fluent
		TableName string
		Statement string
		QB        qb.QueryBuilder
	}

	Cond map[string]any
)

func (f *Fluent) Collection(name string) *Collection {
	col := new(Collection)
	col.Fluent = f
	col.TableName = name

	return col
}

func (c *Collection) Insert(item map[string]any) (*Collection, error) {
	keys := make([]string, 0, len(item))
	values := make([]any, 0, len(item))
	for k, v := range item {
		keys = append(keys, k)
		values = append(values, v)
	}

	c.QB = *new(qb.QueryBuilder)
	c.QB.InsertInto(c.TableName).Columns(keys...).Values(values...)

	return c, nil
}

func (c *Collection) Run() error {
	_, err := c.Fluent.DB.Exec(context.Background(), c.QB.Final, c.QB.InsertValues)

	return err
}

func (c *Collection) Find(cond Cond) (*Collection, error) {
	if len(cond) == 0 {
		return nil, errors.New("condition cannot be empty")
	}

	var field string
	var operator qb.Operator
	var value any

	keys := make([]string, 0, len(cond))
	for k := range cond {
		keys = append(keys, k)
	}

	if len(keys) > 1 {
		// TODO  handle multiple elements later
	}

	if len(keys) == 1 {

		key := keys[0]

		f, o, error := extractOperator(key)

		if error != nil {
			return nil, error
		}

		field = f
		operator = o
		value = cond[key]

	}

	c.QB = *new(qb.QueryBuilder)
	c.QB.SelectAllFrom(c.TableName).Where(field, operator, value)

	return c, nil
}

func (c *Collection) And(cond Cond) *Collection {

	return c
}

func (c *Collection) Or(cond Cond) *Collection {

	return c
}

func (c *Collection) Build() string {
	err := c.QB.Build()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return c.QB.Final
}

func (c *Collection) FetchOne() {

}

func extractOperator(CondStr string) (string, qb.Operator, error) {
	parts := strings.SplitAfterN(CondStr, " ", 2)
	field := parts[0]
	opStr := parts[1]

	// trim leading and trailing space
	trimmedStr := strings.TrimSpace(opStr)
	subStr := strings.Join(strings.Fields(trimmedStr), " ")

	trimmedField := strings.TrimSpace(field)
	field = strings.Join(strings.Fields(trimmedField), "")

	switch subStr {
	case "=":
		return field, qb.EQ, nil
	case "!=":
		return field, qb.NEQ, nil
	case "<>":
		return field, qb.NEQ, nil
	case "<=":
		return field, qb.LTEQ, nil
	case ">=":
		return field, qb.GTEQ, nil
	case ">":
		return field, qb.GT, nil
	case "<":
		return field, qb.LT, nil
	case "NOT IN":
		return field, qb.NotIn, nil
	case "BETWEEN":
		return field, qb.Between, nil
	case "NOT BETWEEN":
		return field, qb.NotBetween, nil
	case "IN":
		return field, qb.In, nil
	case "SIMILAR TO":
		return field, qb.SimilarTo, nil
	case "LIKE":
		return field, qb.Like, nil
	case "ILIKE":
		return field, qb.ILike, nil
	case "IS NULL":
		return field, qb.IsNull, nil
	case "IS NOT NULL":
		return field, qb.IsNotNull, nil
	case "not in":
		return field, qb.NotIn, nil
	case "between":
		return field, qb.Between, nil
	case "not between":
		return field, qb.NotBetween, nil
	case "in":
		return field, qb.In, nil
	case "similar to":
		return field, qb.SimilarTo, nil
	case "like":
		return field, qb.Like, nil
	case "ilike":
		return field, qb.ILike, nil
	case "is null":
		return field, qb.IsNull, nil
	case "is not null":
		return field, qb.IsNotNull, nil
	default:
		return "", "", fmt.Errorf("%s : invalid postgresql operator verify syntax", opStr)
	}
}
