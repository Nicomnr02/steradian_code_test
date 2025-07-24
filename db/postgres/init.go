package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", "postgres", "sec", "localhost", "5430", "steradian")
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse db configuration")
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatal("failed to setup db configuration")
	}

	ctimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := pool.Ping(ctimeout); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected dsn " + dsn)

	return pool
}
