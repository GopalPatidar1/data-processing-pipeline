package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB() *pgxpool.Pool {

	pool, err := pgxpool.New(
		context.Background(),
		"postgres://postgres:mindfire@localhost:5432/pipeline_db",
	)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	DB = pool

	log.Println("Database connected successfully")

	return pool
}
