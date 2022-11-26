package config

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	Port int    `env:"PORT"`
	Dsn  string `env:"MYSQL_DSN"`
}

func LoadEnv(ctx context.Context) (*Env, error) {
	var goenv Env
	err := envconfig.Process(ctx, &goenv)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	return &goenv, nil
}
