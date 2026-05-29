package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:mindfire@localhost:5432/pipeline_db",
	)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	DB = conn

	log.Println("Database connected successfully")
	return conn
}
