// Package routes contains all the handlers plus the http server to run them here
package routes

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/oussamarouabah/template/internal/middlewares"
)

// Store interface sets the required methods from database layer for this pkg
type Store interface {
	GetDummyData(ctx context.Context) error
}

type Server interface {
	Start(ctx context.Context) error
}

type Config struct {
	Port string `mapstructure:"port"`
}

func New(cfg *Config, store Store) Server {
	return &server{
		mux:   echo.New(),
		port:  cfg.Port,
		store: store,
	}
}

type server struct {
	mux   *echo.Echo
	port  string
	store Store
}

func (s *server) Start(ctx context.Context) error {
	s.registerMiddlewares()
	s.setRoutes()
	return s.mux.Start(s.port)
}

func (s *server) registerMiddlewares() {
	// TODO add global middlewares here
	s.mux.Use(middlewares.DummyMiddleware())
}

func (s *server) setRoutes() {
	s.mux.GET("/", dummy, middlewares.SingleMiddleware())
}
