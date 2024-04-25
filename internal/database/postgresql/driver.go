package postgresql

import (
	"context"
	"fmt"
	"os"
	"tryhardplayoffs/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Driver struct {
	Config config.DatabaseConfig
	DSN    string
	Pool   *pgxpool.Pool
}

func (d *Driver) SetDSN() error {
	d.DSN = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		d.Config.User,
		d.Config.Password,
		d.Config.Hostname,
		d.Config.Port,
		d.Config.Name,
	)

	return nil
}

func (d *Driver) NewPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), d.DSN)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return pool, nil
}
