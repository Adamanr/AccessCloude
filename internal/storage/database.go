package storage

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
	Salt string
}

func NewDatabase(url string) *Database {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		slog.Error("Unable to connect to database: " + err.Error())
		os.Exit(1)
	}

	return &Database{
		Conn: conn,
	}
}
