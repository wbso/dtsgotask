package store

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	// ping the database to make sure it's up
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
