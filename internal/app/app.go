package app

import "context"

type Config struct{}

type Application struct {
	done chan struct{}
}

func New(cfg *Config) *Application {
	return &Application{}
}

func (a *Application) Start(ctx context.Context) {}

func (a *Application) Done() chan struct{} {
	return a.done
}
