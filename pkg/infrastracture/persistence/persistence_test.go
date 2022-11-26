package persistence

import (
	"context"
	"database/sql"
	"golang-grpc-sqlboiler-mysql/pkg/config"
	"log"

	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	repo *Repositories
)

func TestMain(m *testing.M) {
	// before
	c := context.Background()
	cfg, err := config.LoadEnv(c)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	repo, err = NewRepositories(c, db)
	if err != nil {
		log.Fatal(err)
	}

	res := m.Run()
	// after

	os.Exit(res)
}
