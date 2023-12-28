package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetDatabaseConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	return conn
}
