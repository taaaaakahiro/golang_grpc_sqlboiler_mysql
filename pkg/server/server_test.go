package server

import (
	"context"
	"database/sql"
	"golang-grpc-sqlboiler-mysql/pkg/config"
	"golang-grpc-sqlboiler-mysql/pkg/infrastracture/persistence"
	"os"
	"testing"
)

var (
	testServer *Server
)

func TestMain(m *testing.M) {
	// before
	c := context.Background()
	cfg, _ := config.LoadEnv(c)
	db, _ := sql.Open("mysql", cfg.Dsn)
	r, _ := persistence.NewRepositories(c, db)
	testServer, _ = NewServer(r)

	res := m.Run()
	// after

	os.Exit(res)
}
