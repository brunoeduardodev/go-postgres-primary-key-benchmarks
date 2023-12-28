package main

import (
	"fmt"

	"github.com/brunoeduardodev/go-postgres-primary-key-benchmarks/internal/seeding"
	"github.com/brunoeduardodev/go-postgres-primary-key-benchmarks/internal/utils"
)

func main() {
	utils.LoadEnvs()
	conn := utils.GetDatabaseConnection()

	fmt.Printf("Connected successfully to database\n")

	seeding.SeedDatabase(conn)

	fmt.Printf("Seeding completed.\n")
}
