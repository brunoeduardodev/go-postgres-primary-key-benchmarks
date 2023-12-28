package main

import (
	"context"
	"fmt"

	"github.com/brunoeduardodev/go-postgres-primary-key-benchmarks/internal/utils"
)

func main() {
	utils.LoadEnvs()
	conn := utils.GetDatabaseConnection()

	fmt.Printf("Connected successfully to database\n")

	defer conn.Close(context.Background())
}
