package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

const postgresDBname = "govision_article"

func NewDB() *DB {
	// TODO: psql Connection
	pool, err := pgxpool.New(context.Background(), postgresDBname)
	if err != nil {
		return nil
	}
	return &DB{
		Pool: pool,
	}
}
