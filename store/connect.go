package store

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MinConns = 2
	config.MaxConns = 10

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// ping the database to make sure it's up
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
