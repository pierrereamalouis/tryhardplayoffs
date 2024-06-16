package handlers

import (
	"tryhardplayoffs/internal/services"

	"github.com/labstack/echo/v4"
)

func AddNhlTeam(ctx echo.Context) error {
	ctn := ctx.Get("container").(*services.Container)

	dbPool := ctn.Pool
}
