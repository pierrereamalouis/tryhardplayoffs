package middleware

import (
	"tryhardplayoffs/internal/services"

	"github.com/labstack/echo/v4"
)

func ContainerMiddleware(ctn *services.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("container", ctn) // store in context
			return next(ctx)
		}
	}
}
