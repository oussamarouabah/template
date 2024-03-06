package app

import (
	"context"

	"github.com/oussamarouabah/template/internal/routes"
	"github.com/oussamarouabah/template/internal/store"
)

type Config struct {
	DatabaseConfig *store.Config  `mapstructure:"mysql"`
	RoutesConfig   *routes.Config `mapstructure:"server"`
}

type Application struct {
	done   chan struct{}
	server routes.Server
	db     *store.Store
}

func New(cfg *Config) *Application {
	db := store.New(cfg.DatabaseConfig)
	return &Application{
		server: routes.New(cfg.RoutesConfig, db),
		db:     db,
	}
}

func (a *Application) Start(ctx context.Context) {

}

func (a *Application) Done() chan struct{} {
	return a.done
}
