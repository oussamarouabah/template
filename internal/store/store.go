package store

import "context"

type Config struct {
	DSN string `mapstructure:"dsn"`
}

type Store struct{}

// New returns a new db layer handler
func New(cfg *Config) *Store { return &Store{} }

func (d *Store) GetDummyData(ctx context.Context) error { return nil }
