package postgresql

import (
	"context"
	"fmt"
	"os"
	"tryhardplayoffs/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Session struct {
	Config config.DatabaseConfig
	DSN    string
	Pool   *pgxpool.Pool
}

func (s *Session) SetDSN() error {
	s.DSN = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		s.Config.User,
		s.Config.Password,
		s.Config.Hostname,
		s.Config.Port,
		s.Config.Name,
	)

	return nil
}

func (s *Session) NewPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), s.DSN)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return pool, nil
}
