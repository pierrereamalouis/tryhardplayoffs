package services

import (
	"fmt"
	"tryhardplayoffs/internal/config"
	"tryhardplayoffs/internal/database/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Pool     *pgxpool.Pool
	Config   config.Config
	Database *postgresql.Session
}

func (c *Container) Shutdown() {
	c.Pool.Close()
}

func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initDatabase()

	return c
}

func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = cfg
}

func (c *Container) initDatabase() {
	sess := postgresql.Session{
		Config: c.Config.Database,
	}

	sess.SetDSN()

	pool, err := sess.NewPool()

	if err != nil {

	}

	c.Pool = pool
}
