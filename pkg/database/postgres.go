package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Unable to parse DatabaseURL: %v\n", err)
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database: %v\n", err)
		pool.Close()
		return nil, err
	}
	return pool, nil

}
