package seeding

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const createSeedsTableQuery = `CREATE TABLE IF NOT EXISTS seeds (
  id SERIAL PRIMARY KEY,
	version INTEGER NOT NULL,
	seeded_at TIMESTAMP NOT NULL
)`

func ensureSeedTableExists(conn *pgx.Conn) {
	_, err := conn.Query(context.Background(), createSeedsTableQuery)

	if err != nil {
		panic(err)
	}
}

const getLatestSeedVersionQuery = `SELECT version FROM seeds LIMIT 1 ORDER BY version DESC`

func getLatestSeedVersion(conn *pgx.Conn) int {
	var version = -1
	err := conn.QueryRow(context.Background(), getLatestSeedVersionQuery).Scan(&version)
	if err != nil {
		return -1
	}

	return version
}

var currentSeedVersion = 1

func checkIfDatabaseIsUpToDate(conn *pgx.Conn) bool {
	latestSeed := getLatestSeedVersion(conn)
	return currentSeedVersion <= latestSeed
}

func SeedDatabase(conn *pgx.Conn) {
	ensureSeedTableExists(conn)
	isDatabaseUpToDate := checkIfDatabaseIsUpToDate(conn)
	if isDatabaseUpToDate {
		return
	}

	fmt.Printf("")

}
