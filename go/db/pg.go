package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

const (
	username = ""
	password = ""
	hostname = ""
	port     = ""
	database = ""
)

// postgresql://<username>:<password>@<hostname>:<port>/<database>

func NewDB() *DB {
	// TODO: psql Connection

	DATABASE_URL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, hostname, port, database)
	pool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		return nil
	}
	return &DB{
		Pool: pool,
	}

}
