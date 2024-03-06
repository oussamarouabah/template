// Package middlewares contains all the middlewares that we need in our app
// These middlwares can be used individually or globaly
package middlewares

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

// DummyMiddleware that logs the request with the time required to handle it
func DummyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			slog.Info("new request", "method", c.Request().Method, "path", c.Path())
			start := time.Now()
			next(c)
			collapsed := time.Since(start)
			slog.Info("request done", "time", collapsed.String())
			return nil
		}
	}
}

func SingleMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			slog.Info("new request", "method", c.Request().Method, "path", c.Path())
			return next(c)
		}
	}
}
