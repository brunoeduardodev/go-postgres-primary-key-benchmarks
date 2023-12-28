package seeding

import (
	"context"
	"fmt"

	fakedata "github.com/brunoeduardodev/go-postgres-primary-key-benchmarks/internal/fake-data"
	"github.com/jackc/pgx/v5"
)

const createUsersTableQuery = `CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	birthdate DATE NOT NULL
)`

func ensureUsersTableExist(conn *pgx.Conn) {
	_, err := conn.Query(context.Background(), createUsersTableQuery)
	if err != nil {
		panic(err)
	}
}

func seedUsers(conn *pgx.Conn, amount int) {
	fakeUsers := fakedata.GenerateFakeUsers(amount)

	fmt.Printf("Fake users: %v\n", fakeUsers)

}
