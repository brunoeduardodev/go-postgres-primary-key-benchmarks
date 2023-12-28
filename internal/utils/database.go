package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDatabaseConnection() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	return conn
}
