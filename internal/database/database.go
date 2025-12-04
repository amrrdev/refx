package database

import (
	"context"

	"github.com/amrrdev/refx/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Queries *db.Queries
	Pool    *pgxpool.Pool
}

func NewDatabase(ctx context.Context, databaseUrl string) (*Database, error) {
	pool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}

	queries := db.New(pool)
	return &Database{
		Pool:    pool,
		Queries: queries,
	}, nil
}
