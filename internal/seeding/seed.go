package seeding

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const createSeedsTableQuery = `CREATE TABLE IF NOT EXISTS seeds (
  id SERIAL PRIMARY KEY,
	version INTEGER NOT NULL,
	seeded_at TIMESTAMP NOT NULL
)`

func ensureSeedTableExists(conn *pgxpool.Pool) {
	_, err := conn.Query(context.Background(), createSeedsTableQuery)

	if err != nil {
		panic(err)
	}
}

const getLatestSeedVersionQuery = `SELECT version FROM seeds LIMIT 1 ORDER BY version DESC`

func getLatestSeedVersion(conn *pgxpool.Pool) int {
	var version = -1
	err := conn.QueryRow(context.Background(), getLatestSeedVersionQuery).Scan(&version)
	if err != nil {
		return -1
	}

	return version
}

var currentSeedVersion = 1

func checkIfDatabaseIsUpToDate(conn *pgxpool.Pool) bool {
	latestSeed := getLatestSeedVersion(conn)
	return currentSeedVersion <= latestSeed
}

func SeedDatabase(conn *pgxpool.Pool) {
	ensureSeedTableExists(conn)
	isDatabaseUpToDate := checkIfDatabaseIsUpToDate(conn)
	if isDatabaseUpToDate {
		return
	}

	ensureUsersTableExist(conn)
	seedUsers(conn, 5)
}
