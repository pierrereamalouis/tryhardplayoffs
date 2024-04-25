package services

import (
	"fmt"
	"tryhardplayoffs/internal/config"
	"tryhardplayoffs/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Pool   *pgxpool.Pool
	Config config.Config
	DAL    *database.DataAccessLayer
}

func (c *Container) Shutdown() {
	c.Pool.Close()
}

func NewContainer() *Container {
	c := new(Container)
	c.initConfig()

	return c
}

func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = cfg
}
