package main

import (
	"os"

	"github.com/florian-renfer/beanbase.io/internal/infrastructure/database"
)

func main() {
	_, err := database.NewDatabaseSQLFactory(database.InstancePostgres)
	if err != nil {
		os.Exit(1)
	}
}
