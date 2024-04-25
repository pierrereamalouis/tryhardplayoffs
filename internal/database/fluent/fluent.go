package fluent

import "github.com/jackc/pgx/v5/pgxpool"

type (
	Fluent struct {
		DB         *pgxpool.Pool
		Collection Collection
	}

	Collection struct {
		Flent     *Fluent
		TableName string
	}
)
