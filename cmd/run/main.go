package main

import (
	"context"
	"fmt"

	"github.com/brunoeduardodev/go-postgres-primary-key-benchmarks/internal"
)

func main() {
	internal.LoadEnvs()
	conn := internal.GetDatabaseConnection()

	fmt.Printf("Connected successfully to database\n")

	defer conn.Close(context.Background())
}
